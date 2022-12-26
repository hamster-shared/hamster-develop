package service

import (
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"gorm.io/gorm"
)

type ILoginService interface {
	GetWorkflowList(workflowType, page, size int) (*[]db2.Workflow, error)
}

type LoginService struct {
	db *gorm.DB
}
