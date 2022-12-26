package service

import (
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"gorm.io/gorm"
)

type IWorkflowService interface {
	GetWorkflowList(workflowType, page, size int) (*[]db2.Workflow, error)
}

type WorkflowService struct {
	db *gorm.DB
}

func NewWorkflowService() *WorkflowService {
	return &WorkflowService{}
}

func (w *WorkflowService) Init(db *gorm.DB) {
	w.db = db
}

func (w *WorkflowService) GetWorkflowList(workflowType, page, size int) (*[]db2.Workflow, error) {
	return nil, nil
}
