package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService() *ReportService {
	return &ReportService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (c *ReportService) QueryReports(projectId, reportType string, Type string, page int, size int) (vo.Page[db2.Report], error) {
	var total int64
	var reports []db2.Report
	tx := c.db.Model(db2.Report{}).Where("project_id = ?", projectId)
	if Type != "" || reportType != "" {
		tx = tx.Where("check_tool = ? and name=?", Type, reportType)
	} else if Type == "" && reportType == "" {
		tx = tx.Where("1 = 1")
	} else if Type != "" {
		tx = tx.Where("check_tool = ?", Type)
	} else {
		tx = tx.Where("name = ?", reportType)
	}
	tx = tx.Order("create_time desc")
	result := tx.Offset((page - 1) * size).Limit(size).Find(&reports).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Report](), result.Error
	}

	return vo.NewPage[db2.Report](reports, int(total), page, size), nil
}

func (c *ReportService) QueryFrontendReports(projectId string, page int, size int) (vo.Page[db2.Report], error) {
	var total int64
	var reports []db2.Report
	tx := c.db.Model(db2.Report{}).Where("project_id = ?", projectId)
	result := tx.Offset((page - 1) * size).Limit(size).Find(&reports).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Report](), result.Error
	}
	return vo.NewPage[db2.Report](reports, int(total), page, size), nil
}

func (c *ReportService) QueryReportsByWorkflow(workflowId, workflowDetailId int) ([]vo.ReportVo, error) {
	var reports []db2.Report
	var result []vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res.Error != nil {
		return result, res.Error
	}
	if len(reports) > 0 {
		copier.Copy(&result, &reports)
	}
	return result, nil
}

func (c *ReportService) QueryFrontendReportsByWorkflow(workflowId, workflowDetailId int) ([]vo.ReportVo, error) {
	var reports []db2.Report
	var data []vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res.Error != nil {
		return data, res.Error
	}
	_ = copier.Copy(&data, &reports)
	return data, nil
}

func (c *ReportService) QueryReportCheckTools(projectId string) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Report{}).Distinct("check_tool").Select("check_tool").Where("project_id = ?", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}
