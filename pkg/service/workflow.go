package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/engine"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
	"gorm.io/gorm"
)

type WorkflowService struct {
	db     *gorm.DB
	engine *engine.Engine
}

func NewWorkflowService() *WorkflowService {
	return &WorkflowService{
		db:     application.GetBean[*gorm.DB]("db"),
		engine: application.GetBean[*engine.Engine]("engine"),
	}
}

func (w *WorkflowService) ExecProjectCheckWorkflow(projectId uint, user vo.UserAuth) error {
	return w.ExecProjectWorkflow(projectId, user, 1)
}

func (w *WorkflowService) ExecProjectBuildWorkflow(projectId uint, user vo.UserAuth) error {
	return w.ExecProjectWorkflow(projectId, user, 2)
}

func (w *WorkflowService) ExecProjectWorkflow(projectId uint, user vo.UserAuth, workflowType uint) error {

	// query project workflow

	var workflow db.Workflow

	w.db.Where(&db.Workflow{
		ProjectId: projectId,
		Type:      workflowType,
	}).First(&workflow)

	if &workflow == nil {
		return errors.New("no check workflow in the project ")
	}

	workflowKey := GetWorkflowKey(projectId, workflow.Id)

	detail, err := w.engine.ExecuteJob(workflowKey)

	if err != nil {
		return err
	}
	stageInfo, err := json.Marshal(detail.Stages)
	if err != nil {
		return err
	}

	dbDetail := db.WorkflowDetail{
		WorkflowId:  workflow.Id,
		ExecNumber:  uint(detail.Id),
		StageInfo:   string(stageInfo),
		TriggerUser: user.Username,
		TriggerMode: 1,
		CodeInfo:    "",
		Status:      uint(detail.Status),
		StartTime: sql.NullTime{
			Time:  detail.StartTime,
			Valid: true,
		},
	}

	err = w.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(&dbDetail)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkflowService) GetWorkflowList(workflowType, page, size int) (*[]db.Workflow, error) {
	return nil, nil
}

func GetWorkflowKey(projectId uint, workflowId uint) string {
	return fmt.Sprintf("%d_%d", projectId, workflowId)
}
