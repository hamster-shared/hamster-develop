package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"gorm.io/gorm"
)

type FrontendPackageService struct {
	db *gorm.DB
}

func NewFrontendPackageService() *FrontendPackageService {
	return &FrontendPackageService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (f *FrontendPackageService) QueryFrontendPackages(projectId string, page int, size int) (vo.Page[db2.FrontendPackage], error) {
	var total int64
	var packages []db2.FrontendPackage
	tx := f.db.Model(db2.FrontendPackage{}).Where("project_id = ?", projectId)
	result := tx.Offset((page - 1) * size).Limit(size).Find(&packages).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.FrontendPackage](), result.Error
	}

	return vo.NewPage[db2.FrontendPackage](packages, int(total), page, size), nil
}

func (f *FrontendPackageService) QueryPackageById(id int) (db2.FrontendPackage, error) {
	var packages db2.FrontendPackage
	res := f.db.Model(db2.FrontendPackage{}).Where("id = ?", id).First(&packages).Error
	if res != nil {
		return packages, res
	}
	return packages, nil
}

func (f *FrontendPackageService) QueryPackages(workflowId, workflowDetailId int) ([]db2.FrontendPackage, error) {
	var packages []db2.FrontendPackage
	res := f.db.Model(db2.FrontendPackage{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&packages).Error
	if res != nil {
		return packages, res
	}
	return packages, nil
}

func (f *FrontendPackageService) UpdateFrontedPackage(data db2.FrontendPackage) error {
	return f.db.Save(&data).Error
}

func (f *FrontendPackageService) QueryFrontendDeployInfo(workflowId, workflowDetailId int) (db2.FrontendDeploy, error) {
	var packageDeploy db2.FrontendDeploy
	res := f.db.Model(db2.FrontendDeploy{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).First(&packageDeploy).Error
	if res != nil {
		return packageDeploy, res
	}
	return packageDeploy, nil
}

func (f *FrontendPackageService) QueryFrontendDeployById(id int) (db2.FrontendDeploy, error) {
	var packageDeploy db2.FrontendDeploy
	res := f.db.Model(db2.FrontendDeploy{}).Where("package_id = ?", id).First(&packageDeploy).Error
	if res != nil {
		return packageDeploy, res
	}
	return packageDeploy, nil
}

func (f *FrontendPackageService) DeleteFrontendDeploy(id int) error {
	//err := f.db.Debug().Where("id = ?", workflowDetailId).Delete(&db2.WorkflowDetail{}).Error
	//if err != nil {
	//	return err
	//}
	err := f.db.Debug().Where("package_id = ? ", id).Delete(&db2.FrontendDeploy{}).Error
	return err
}
