package service

import (
	"github.com/hamster-shared/a-line/pkg/application"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
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

func (c *ContractService) QueryContracts(projectId uint, query, version string, page int, size int) (vo.Page[db2.Contract], error) {
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

	result := tx.Offset((page - 1) * size).Limit(size).Find(&contracts).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Contract](), result.Error
	}

	return vo.NewPage[db2.Contract](contracts, int(total), page, size), nil
}
