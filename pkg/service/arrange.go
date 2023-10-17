package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"gorm.io/gorm"
	"strconv"
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
	arrange := db2.ContractArrange{
		ProjectId:       projectId,
		Version:         param.Version,
		OriginalArrange: param.OriginalArrange,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	err = a.db.Model(db2.ContractArrange{}).Create(&arrange).Error
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(arrange.Id)), nil
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
		FkArrangeId:        strconv.Itoa(int(contractArrange.Id)),
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
	err = a.db.Raw("SELECT cd.id, c.name AS contractName, cd.contract_id, cd.project_id, cd.version, cd.deploy_time, cd.network, cd.address, cd.type, cd.declare_tx_hash, cd.deploy_time, cd.status, cd.abi_info from t_contract_deploy cd LEFT JOIN t_contract c ON cd.contract_id = c.id WHERE cd.project_id = ? AND cd.version = ?", projectId, version).Scan(&deployContractList).Error
	if err != nil {
		return list, err
	}
	location := time.FixedZone("GMT+8", 8*60*60)
	for _, deployContract := range deployContractList {
		deployContract.DeployTime = deployContract.DeployTime.In(location)
		deployContract.DeployTimeFormat = deployContract.DeployTime.Format("02.Jan.2006 15:04 MST")
	}
	return deployContractList, nil
}
