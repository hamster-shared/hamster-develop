package service

import (
	"encoding/json"
	"fmt"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/samber/lo"
)

type ArrangeProcessData struct {
	DeployStep []ArrangeDeployStep `json:"deployStep"`
	Step       int                 `json:"step"`
}

type ArrangeDeployStep struct {
	Contract ArrangeContract `json:"contract"`
	Steps    []ArrangeStep   `json:"steps"`
	Step     int             `json:"step"`
	Status   string          `json:"status"`
}

type ArrangeContract struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	ProxyAddress string `json:"proxyAddress"`
	Proxy        bool   `json:"proxy"`
}

type ArrangeStep struct {
	Type      string   `json:"type"`
	Method    string   `json:"method"`
	Params    []string `json:"params"`
	Status    string   `json:"status"` // SUCCESS, PENDDING ,FAILED
	ErrorInfo string   `json:"errorInfo"`
}

func UnmarshalProcessData(processData string) (ArrangeProcessData, error) {
	var obj ArrangeProcessData

	err := json.Unmarshal([]byte(processData), &obj)
	if err != nil {
		fmt.Println(err)
		return ArrangeProcessData{}, err
	}

	steps := lo.FilterMap(obj.DeployStep, func(item ArrangeDeployStep, _ int) (ArrangeDeployStep, bool) {
		return item, item.Contract.Name != ""
	})
	obj.DeployStep = steps
	return obj, nil
}

func (p *ArrangeProcessData) toJobDetail() []model.StageDetail {

	var result []model.StageDetail

	for _, deployStep := range p.DeployStep {

		steps := lo.Map(deployStep.Steps, func(item ArrangeStep, index int) model.Step {
			name := "constructor"
			if item.Method == "" {
				name = item.Type
			} else {
				name = item.Method
			}
			step := model.Step{
				Name:     name,
				Status:   toEngineStatus(item.Status),
				Duration: 0,
			}

			return step
		})
		stage := model.Stage{
			Steps: steps,
		}
		stageDetail := model.NewStageDetail(deployStep.Contract.Name, stage)
		stageDetail.Status = toEngineStatus(deployStep.Status)
		result = append(result, stageDetail)
	}
	return result
}

func (p *ArrangeProcessData) toJobDetailString() string {
	data := p.toJobDetail()
	marshal, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func toEngineStatus(status string) model.Status {
	if status == "SUCCESS" {
		return model.STATUS_SUCCESS
	} else if status == "PENDDING" {
		return model.STATUS_NOTRUN
	} else if status == "FAILED" {
		return model.STATUS_FAIL
	} else if status == "STOP" {
		return model.STATUS_STOP
	}

	return model.STATUS_FAIL
}

func (p *ArrangeProcessData) GetStatus() model.Status {

	if len(p.DeployStep) == 0 {
		return model.STATUS_FAIL
	}

	status := p.DeployStep[0].Status

	for _, ds := range p.DeployStep {
		if ds.Status != status {
			return toEngineStatus(ds.Status)
		}
	}

	return toEngineStatus(status)
}

func (p *ArrangeProcessData) GetErrorInfo() string {
	for _, ds := range p.DeployStep {
		if ds.Status == "FAILED" {
			for _, s := range ds.Steps {
				if s.Status == "FAILED" {
					return s.ErrorInfo
				}
			}
		}
	}
	return ""
}
