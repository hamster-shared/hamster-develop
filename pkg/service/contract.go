package service

import (
	"github.com/goperate/convert/core/array"
	"github.com/hamster-shared/a-line/pkg/application"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ContractService struct {
	db *gorm.DB
}

func NewContractService() *ContractService {
	return &ContractService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (c *ContractService) SaveDeploy(entity db2.ContractDeploy) error {

	err := c.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(entity)
		return nil
	})

	return err
}

func (c *ContractService) QueryContracts(projectId uint, query, version, network string, page int, size int) (vo.Page[db2.Contract], error) {
	var total int64
	var contracts []db2.Contract
	tx := c.db.Model(db2.Contract{
		ProjectId: projectId,
	})
	if query != "" {
		tx = tx.Where("name like ? ", "%"+query+"%")
	}

	if version != "" {
		tx = tx.Where("version = ?", version)
	}
	if network != "" {
		tx = tx.Where("network like ?", "%"+network+"%")
	}

	result := tx.Offset((page - 1) * size).Limit(size).Find(&contracts).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Contract](), result.Error
	}

	return vo.NewPage[db2.Contract](contracts, int(total), page, size), nil
}

func (c *ContractService) QueryContractByWorkflow(workflowId, workflowDetailId int) ([]db2.Contract, error) {
	var contracts []db2.Contract
	res := c.db.Model(db2.Contract{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&contracts)
	if res != nil {
		return contracts, res.Error
	}
	return contracts, nil
}

func (c *ContractService) QueryContractByVersion(projectId int, version string) ([]vo.ContractVo, error) {
	var contracts []db2.Contract
	var data []vo.ContractVo
	res := c.db.Model(db2.Contract{}).Where("project_id = ? and version = ?", projectId, version).Find(&contracts)
	if res != nil {
		return data, res.Error
	}
	if len(contracts) > 0 {
		copier.Copy(&data, &contracts)
	}
	return data, nil
}

func (c *ContractService) QueryContractDeployByVersion(projectId int, version string) (vo.ContractDeployInfoVo, error) {
	var data vo.ContractDeployInfoVo
	var contractDeployData []db2.ContractDeploy
	res := c.db.Model(db2.ContractDeploy{}).Where("project_id = ? and version = ?", projectId, version).Find(&contractDeployData)
	if res.Error != nil {
		return data, res.Error
	}
	contractInfo := make(map[string]vo.ContractInfoVo)
	if len(contractDeployData) > 0 {
		arr := array.NewObjArray(contractDeployData, "ContractId")
		res2 := arr.ToIdMapArray().(map[uint][]db2.ContractDeploy)
		for u, deploys := range res2 {
			var contractData db2.Contract
			res := c.db.Model(db2.Contract{}).Where("id = ?", u).First(&contractData)
			if res.Error == nil {
				var deployInfo []vo.DeployInfVo
				if len(deploys) > 0 {
					for _, deploy := range deploys {
						var deployData vo.DeployInfVo
						copier.Copy(&deployData, &deploy)
						deployInfo = append(deployInfo, deployData)
					}
				}
				var contractInfoVo vo.ContractInfoVo
				copier.Copy(&contractInfoVo, &contractData)
				contractInfoVo.DeployInfo = deployInfo
				contractInfo[contractData.Name] = contractInfoVo
			}
		}
	}
	data.Version = version
	data.ContractInfo = contractInfo
	return data, nil
}

func (c *ContractService) QueryVersionList(projectId int) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Contract{}).Distinct("version").Select("version").Where("project_id = ?", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}
