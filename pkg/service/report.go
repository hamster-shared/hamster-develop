package service

import (
	"github.com/hamster-shared/a-line/pkg/application"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
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

func (c *ReportService) QueryReports(projectId uint, Type uint, page int, size int) (vo.Page[db2.Report], error) {
	var total int64
	var reports []db2.Report
	tx := c.db.Model(db2.Report{
		ProjectId: projectId,
	})
	if Type != 0 {
		tx = tx.Where("type = ?", Type)
	}
	result := tx.Offset((page - 1) * size).Limit(size).Find(&reports).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Report](), result.Error
	}

	return vo.NewPage[db2.Report](reports, int(total), page, size), nil
}

func (c *ReportService) QueryReportsByWorkflow(workflowId, workflowDetailId int) ([]vo.ReportVo, error) {
	var reports []db2.Report
	var data []vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res != nil {
		return data, res.Error
	}
	if len(reports) > 0 {
		copier.Copy(&data, &reports)
	}
	return data, nil
}
