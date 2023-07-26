package service

import (
	"database/sql"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type DfxDataService struct {
	db *gorm.DB
}

func NewDfxDataService() *DfxDataService {
	return &DfxDataService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (d *DfxDataService) SaveDfxJsonData(projectId string, jsonData string) error {
	var dfxData db.IcpDfxData
	err := d.db.Where("project_id = ?", projectId).First(&dfxData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			dfxData.ProjectId = projectId
			dfxData.DfxData = jsonData
			dfxData.CreateTime = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			err = d.db.Create(&dfxData).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}
func (d *DfxDataService) QueryDfxJsonDataByProjectId(projectId string) (vo.IcpDfxDataVo, error) {
	var data db.IcpDfxData
	var vo vo.IcpDfxDataVo
	err := d.db.Model(db.IcpDfxData{}).Where("project_id = ?", projectId).First(&data).Error
	if err != nil {
		return vo, err
	}
	copier.Copy(&vo, &data)
	return vo, nil
}
func (d *DfxDataService) UpdateDfxJsonData(id int, jsonData string) error {
	var data db.IcpDfxData
	err := d.db.Model(db.IcpDfxData{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return err
	}
	data.DfxData = jsonData
	err = d.db.Save(&data).Error
	return err
}
