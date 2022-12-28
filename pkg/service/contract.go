package service

import (
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
