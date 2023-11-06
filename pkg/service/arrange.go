package service

import (
	"encoding/json"
	"errors"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type ArrangeService struct {
	db *gorm.DB
}

func NewArrangeService() *ArrangeService {
	return &ArrangeService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (a *ArrangeService) GetToBeArrangedContractList(id, version string) (vo.ToBeArrangedContractListVo, error) {
	var list vo.ToBeArrangedContractListVo
	projectId, err := uuid.FromString(id)
	if err != nil {
		return list, err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return list, err
	}
	var contracts []db2.Contract
	err = a.db.Model(db2.Contract{}).Where("project_id = ? and version = ?", projectId, version).Find(&contracts).Error
	if err != nil {
		return list, err
	}
	if len(contracts) < 1 {
		return list, errors.New("there are no contracts in this version")
	}
	var contractNameArrangeParam parameter.ContractNameArrangeParam
	contractNameArrangeParam.ProjectId = id
	contractNameArrangeParam.Version = version
	var contractNameList []string
	contractMap := make(map[string]db2.Contract)
	for _, contract := range contracts {
		contractNameList = append(contractNameList, contract.Name)
		contractMap[contract.Name] = contract
	}

	//获取编排的合约名
	var arrangeContractName string
	err = a.db.Model(db2.ContractArrange{}).Select("arrange_contract_name").Where("project_id = ?", projectId).Order("update_time desc").Limit(1).First(&arrangeContractName).Error
	if err == nil {
		//存在历史版本
		var contractNameArrange parameter.ContractNameArrange
		err := json.Unmarshal([]byte(arrangeContractName), &contractNameArrange)
		if err != nil {
			return list, err
		}

		//获取该项目的最新的编排信息
		var arrangeCacheList []db2.ContractArrangeCache
		err = a.db.Model(db2.ContractArrangeCache{}).Distinct("contract_name").Select("contract_name", "original_arrange").Where("project_id = ?", projectId).Order("update_time desc").Find(&arrangeCacheList).Error
		if err != nil {
			return list, err
		}
		contractArrangeMap := make(map[string]string)
		for _, cache := range arrangeCacheList {
			contractArrangeMap[cache.ContractName] = cache.OriginalArrange
		}
		var useContract []vo.ToBeArrangedContractVo
		var noUseContract []vo.ToBeArrangedContractVo
		var usedNames []string
		var saveUsedNames []string
		var saveNoUsedNames []string
		for _, name := range contractNameArrange.UseContract {
			contractName := name
			if lastIndex := strings.LastIndex(contractName, "("); lastIndex >= 0 {
				contractName = contractName[:lastIndex]
			}
			if contract, exists := contractMap[contractName]; exists {
				var toBeArrangedContractVo vo.ToBeArrangedContractVo
				copier.Copy(&toBeArrangedContractVo, &contract)
				toBeArrangedContractVo.Name = name
				toBeArrangedContractVo.OriginalArrange = contractArrangeMap[name]
				usedNames = append(usedNames, name)
				saveUsedNames = append(saveUsedNames, name)
				useContract = append(useContract, toBeArrangedContractVo)
			}
		}
		for _, name := range contractNameArrange.NoUseContract {
			contractName := name
			if lastIndex := strings.LastIndex(contractName, "("); lastIndex >= 0 {
				contractName = contractName[:lastIndex]
			}
			if contract, exists := contractMap[contractName]; exists {
				var toBeArrangedContractVo vo.ToBeArrangedContractVo
				copier.Copy(&toBeArrangedContractVo, &contract)
				toBeArrangedContractVo.Name = name
				usedNames = append(usedNames, name)
				saveNoUsedNames = append(saveNoUsedNames, name)
				noUseContract = append(noUseContract, toBeArrangedContractVo)
			}
		}
		//获取没有使用的合约名
		newContractName := utils.Difference(contractNameList, usedNames)
		for _, name := range newContractName {
			if contract, exists := contractMap[name]; exists {
				var toBeArrangedContractVo vo.ToBeArrangedContractVo
				copier.Copy(&toBeArrangedContractVo, &contract)
				toBeArrangedContractVo.Name = name
				toBeArrangedContractVo.OriginalArrange = contractArrangeMap[name]
				saveUsedNames = append(saveUsedNames, name)
				useContract = append(useContract, toBeArrangedContractVo)
			}
		}
		if useContract == nil {
			list.UseContract = []vo.ToBeArrangedContractVo{}
		} else {
			list.UseContract = useContract
		}
		if noUseContract == nil {
			list.NoUseContract = []vo.ToBeArrangedContractVo{}
		} else {
			list.NoUseContract = noUseContract
		}
		contractNameArrangeParam.UseContract = saveUsedNames
		contractNameArrangeParam.NoUseContract = saveNoUsedNames
		a.SaveContractNameArrange(contractNameArrangeParam)
		return list, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//进行初始化
		contractNameArrangeParam.UseContract = contractNameList
		a.SaveContractNameArrange(contractNameArrangeParam)
		var toBeArrangedContractVoList []vo.ToBeArrangedContractVo
		copier.Copy(&toBeArrangedContractVoList, &contracts)
		list.UseContract = toBeArrangedContractVoList
		return list, nil
	} else {
		return list, err
	}
}

func (a *ArrangeService) GetArrangedDataList(id, version string) ([]string, error) {
	var list []string
	projectId, err := uuid.FromString(id)
	if err != nil {
		return list, err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return list, err
	}

	//获取编排的合约名
	var arrangeContractName string
	err = a.db.Model(db2.ContractArrange{}).Select("arrange_contract_name").Where("project_id = ? and version = ?", projectId, version).Order("update_time desc").Limit(1).First(&arrangeContractName).Error
	if err == nil {
		//存在历史版本
		var contractNameArrange parameter.ContractNameArrange
		err := json.Unmarshal([]byte(arrangeContractName), &contractNameArrange)
		if err != nil {
			return list, err
		}
		var ContractArrangeCacheList []db2.ContractArrangeCache
		err = a.db.Model(db2.ContractArrangeCache{}).Where("project_id = ? and version = ?", projectId, version).Find(&ContractArrangeCacheList).Error
		if err != nil {
			return list, err
		}
		contractNameMap := make(map[string]string)
		for _, contractArrangeCache := range ContractArrangeCacheList {
			contractNameMap[contractArrangeCache.ContractName] = contractArrangeCache.OriginalArrange
		}
		for _, name := range contractNameArrange.UseContract {
			if originalArrange, exists := contractNameMap[name]; exists {
				list = append(list, originalArrange)
			} else {
				list = append(list, "")
			}
		}
		return list, nil
	} else {
		return list, nil
	}
}

func (a *ArrangeService) GetOriginalArrangedData(id, version string) (string, error) {
	projectId, err := uuid.FromString(id)
	if err != nil {
		return "", err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return "", err
	}

	var originalArrange string
	err = a.db.Model(db2.ContractArrange{}).Select("original_arrange").Where("project_id = ? and version = ?", projectId, version).First(&originalArrange).Error
	return originalArrange, err
}

func (a *ArrangeService) SaveContractArrange(param parameter.ContractArrangeParam) (string, error) {
	projectId, err := uuid.FromString(param.ProjectId)
	if err != nil {
		return "", err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return "", err
	}
	var contractArrange db2.ContractArrange
	err = a.db.Model(db2.ContractArrange{}).Where("project_id = ? and version = ?", projectId, param.Version).First(&contractArrange).Error
	if err == nil {
		contractArrange.OriginalArrange = param.OriginalArrange
		contractArrange.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrange{}).Where("id = ?", contractArrange.Id).Updates(&contractArrange)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		contractArrange.ProjectId = projectId
		contractArrange.Version = param.Version
		contractArrange.OriginalArrange = param.OriginalArrange
		contractArrange.CreateTime = time.Now()
		contractArrange.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrange{}).Create(&contractArrange)
	} else {
		return "", err
	}
	return strconv.Itoa(int(contractArrange.Id)), nil
}

func (a *ArrangeService) SaveContractNameArrange(param parameter.ContractNameArrangeParam) (string, error) {
	projectId, err := uuid.FromString(param.ProjectId)
	if err != nil {
		return "", err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return "", err
	}
	var contractNameArrange parameter.ContractNameArrange
	contractNameArrange.UseContract = param.UseContract
	contractNameArrange.NoUseContract = param.NoUseContract
	jsonData, err := json.Marshal(contractNameArrange)
	if err != nil {
		return "", err
	}

	var contractArrange db2.ContractArrange
	err = a.db.Model(db2.ContractArrange{}).Where("project_id = ? and version = ?", projectId, param.Version).First(&contractArrange).Error
	if err == nil {
		contractArrange.ArrangeContractName = string(jsonData)
		contractArrange.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrange{}).Where("id = ?", contractArrange.Id).Updates(&contractArrange)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		contractArrange.ProjectId = projectId
		contractArrange.Version = param.Version
		contractArrange.ArrangeContractName = string(jsonData)
		contractArrange.CreateTime = time.Now()
		contractArrange.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrange{}).Create(&contractArrange)
	} else {
		return "", err
	}
	return strconv.Itoa(int(contractArrange.Id)), nil
}

func (a *ArrangeService) UpdateContractArrange(param parameter.ContractArrangeParam) (string, error) {
	projectId, err := uuid.FromString(param.ProjectId)
	if err != nil {
		return "", err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return "", err
	}
	var contractArrange db2.ContractArrange
	err = a.db.Model(db2.ContractArrange{}).Where("project_id = ?", projectId).First(&contractArrange).Error
	if err != nil {
		return "", err
	}
	contractArrange.OriginalArrange = param.OriginalArrange
	if param.Version != "" {
		contractArrange.Version = param.Version
	}
	contractArrange.UpdateTime = time.Now()
	err = a.db.Model(db2.ContractArrange{}).Where("project_id = ?", projectId).Updates(&contractArrange).Error
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(contractArrange.Id)), nil
}

func (a *ArrangeService) SaveContractArrangeExecute(param parameter.ContractArrangeExecuteParam) (string, error) {
	projectId, err := uuid.FromString(param.ProjectId)
	if err != nil {
		return "", err
	}
	var contractArrange db2.ContractArrange
	err = a.db.Model(db2.ContractArrange{}).Where("id = ?", param.FkArrangeId).First(&contractArrange).Error
	if err != nil {
		return "", err
	}
	arrangeExecute := db2.ContractArrangeExecute{
		ProjectId:          projectId,
		FkArrangeId:        contractArrange.Id,
		Version:            param.Version,
		Network:            param.Network,
		ArrangeProcessData: param.ArrangeProcessData,
		CreateTime:         time.Now(),
		UpdateTime:         time.Now(),
	}
	err = a.db.Model(db2.ContractArrangeExecute{}).Create(&arrangeExecute).Error
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(arrangeExecute.Id)), nil
}

func (a *ArrangeService) UpdateContractArrangeExecute(update parameter.ContractArrangeExecuteUpdateParam) (string, error) {
	var arrangeExecute db2.ContractArrangeExecute
	err := a.db.Model(db2.ContractArrangeExecute{}).Where("id = ?", update.Id).First(&arrangeExecute).Error
	if err != nil {
		return "", err
	}
	arrangeExecute.ArrangeProcessData = update.ArrangeProcessData
	arrangeExecute.UpdateTime = time.Now()
	err = a.db.Model(db2.ContractArrangeExecute{}).Where("id = ?", update.Id).Updates(&arrangeExecute).Error
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(arrangeExecute.Id)), nil
}

func (a *ArrangeService) GetContractArrangeExecuteInfo(executeId string) (info db2.ContractArrangeExecute, err error) {
	var arrangeExecute db2.ContractArrangeExecute
	err = a.db.Model(db2.ContractArrangeExecute{}).Where("id = ?", executeId).First(&arrangeExecute).Error
	if err != nil {
		return info, err
	}
	return arrangeExecute, nil
}

func (a *ArrangeService) GetDeployArrangeContractList(projectId, version string) (list []vo.DeployContractListVo, err error) {
	var deployContractList []vo.DeployContractListVo
	err = a.db.Raw("SELECT cd.id, c.name AS contract_name, cd.contract_id, cd.project_id, cd.version, cd.deploy_time, cd.network, cd.address, cd.type, cd.declare_tx_hash, cd.deploy_time, cd.status, cd.abi_info from t_contract_deploy cd LEFT JOIN t_contract c ON cd.contract_id = c.id WHERE cd.project_id = ? AND cd.version = ?", projectId, version).Scan(&deployContractList).Error
	if err != nil {
		return list, err
	}
	for _, deployContractVo := range deployContractList {
		var deployContract vo.DeployContractListVo
		copier.Copy(&deployContract, &deployContractVo)
		deployContract.DeployTimeFormat = deployContractVo.DeployTime.Format("Jan-02-2006 03:04:05 PM MST")
		list = append(list, deployContract)
	}
	return list, nil
}

func (a *ArrangeService) SaveContractArrangeCache(param parameter.ContractArrangeCacheParam) (vo.ContractArrangeCacheVo, error) {
	var vo vo.ContractArrangeCacheVo
	projectId, err := uuid.FromString(param.ProjectId)
	if err != nil {
		return vo, err
	}
	var project db2.Project
	err = a.db.Model(db2.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return vo, err
	}
	var contractArrangeCache db2.ContractArrangeCache
	err = a.db.Model(db2.ContractArrangeCache{}).Where("project_id = ? and contract_id = ? and contract_name = ? and version = ?", projectId, param.ContractId, param.ContractName, param.Version).First(&contractArrangeCache).Error
	if err == nil {
		contractArrangeCache.OriginalArrange = param.OriginalArrange
		contractArrangeCache.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrangeCache{}).Where("id = ?", contractArrangeCache.Id).Updates(&contractArrangeCache)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		contractArrangeCache.ContractId = param.ContractId
		contractArrangeCache.ProjectId = projectId
		contractArrangeCache.OriginalArrange = param.OriginalArrange
		contractArrangeCache.Version = param.Version
		contractArrangeCache.ContractName = param.ContractName
		contractArrangeCache.CreateTime = time.Now()
		contractArrangeCache.UpdateTime = time.Now()
		a.db.Model(db2.ContractArrangeCache{}).Create(&contractArrangeCache)
	} else {
		return vo, err
	}
	copier.Copy(&vo, &contractArrangeCache)
	return vo, nil
}

func (a *ArrangeService) GetContractArrangeCache(query parameter.ContractArrangeCacheQuery) (string, error) {
	projectId, err := uuid.FromString(query.ProjectId)
	if err != nil {
		return "", err
	}
	var contractArrangeCache db2.ContractArrangeCache
	err = a.db.Model(db2.ContractArrangeCache{}).Where("project_id = ? and contract_name = ?", projectId, query.ContractName).Order("update_time desc").Limit(1).Find(&contractArrangeCache).Error
	if err != nil {
		return "", err
	}
	return contractArrangeCache.OriginalArrange, nil
}

func (a *ArrangeService) GetContractInfo(query parameter.ContractInfoQuery) (vo.ToBeArrangedContractVo, error) {
	var contract db2.Contract
	err := a.db.Model(db2.Contract{}).Where("id = ?", query.Id).First(&contract).Error
	if err != nil {
		return vo.ToBeArrangedContractVo{}, err
	}
	var toBeArrangedContractVo vo.ToBeArrangedContractVo
	copier.Copy(&toBeArrangedContractVo, &contract)
	toBeArrangedContractVo.Name = query.Name
	return toBeArrangedContractVo, nil
}
