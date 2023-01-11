package action

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hamster-shared/a-line/engine/consts"
	"github.com/hamster-shared/a-line/engine/logger"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/engine/output"
	"github.com/hamster-shared/a-line/pkg/utils"
	"os"
	path2 "path"
	"strconv"
	"strings"
)

// CheckAggregationAction 合约聚合
type CheckAggregationAction struct {
	path   string
	ctx    context.Context
	output *output.Output
}

func NewCheckAggregationAction(step model.Step, ctx context.Context, output *output.Output) *CheckAggregationAction {
	return &CheckAggregationAction{
		path:   step.With["path"],
		ctx:    ctx,
		output: output,
	}
}

func (a *CheckAggregationAction) Pre() error {
	return nil
}

func (a *CheckAggregationAction) Hook() (*model.ActionResult, error) {
	a.output.NewStep("check-aggregation")

	stack := a.ctx.Value(STACK).(map[string]interface{})
	jobName, ok := stack["name"].(string)
	if !ok {
		return nil, errors.New("get job name error")
	}
	jobId, ok := stack["id"].(string)
	if !ok {
		return nil, errors.New("get job id error")
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Errorf("Failed to get home directory, the file will be saved to the current directory, err is %s", err.Error())
		userHomeDir = "."
	}

	var absPathList []string
	destDir := path2.Join(userHomeDir, consts.ArtifactoryDir, jobName, consts.CheckName, jobId)
	absPathList = utils.GetSameFileNameFiles(destDir, consts.CheckResult, absPathList)
	_, err = os.Stat(destDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	var contractCheckResultList []model.ContractCheckResult[string]
	var methodsPropertiesReport model.ContractCheckResult[string]
	//var styleGuideValidationsReportList model.ContractCheckResult[[]model.ContractStyleGuideValidationsReportDetails]
	//var securityAnalysisReportList model.ContractCheckResult[[]model.ContractStyleGuideValidationsReportDetails]
	for _, path := range absPathList {
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, errors.New("file open fail")
		}
		result := string(file)
		if strings.Contains(result, consts.ContractMethodsPropertiesReport.Name) {
			err := json.Unmarshal(file, &methodsPropertiesReport)
			if err != nil {
				continue
			}
			contractCheckResultList = append(contractCheckResultList, methodsPropertiesReport)
		}
		if strings.Contains(result, consts.ContractStyleGuideValidationsReport.Name) {
			var styleGuideValidationsReport model.ContractCheckResult[[]model.ContractStyleGuideValidationsReportDetails]
			err := json.Unmarshal(file, &styleGuideValidationsReport)
			if err != nil {
				continue
			}
			var styleGuideValidationsReportString model.ContractCheckResult[string]
			styleGuideValidationsReportString.Tool = styleGuideValidationsReport.Tool
			styleGuideValidationsReportString.Name = styleGuideValidationsReport.Name
			styleGuideValidationsReportString.Result = styleGuideValidationsReport.Result
			var contractCheckResultDetailsList []model.ContractCheckResultDetails[string]
			for _, report := range styleGuideValidationsReport.Context {
				var contractCheckResultDetails model.ContractCheckResultDetails[string]
				contractCheckResultDetails.Name = report.Name
				contractCheckResultDetails.Issue = report.Issue
				marshal, err := json.Marshal(report.Message)
				if err != nil {
					continue
				}
				contractCheckResultDetails.Message = string(marshal)
				contractCheckResultDetailsList = append(contractCheckResultDetailsList, contractCheckResultDetails)
			}
			styleGuideValidationsReportString.Context = contractCheckResultDetailsList
			contractCheckResultList = append(contractCheckResultList, styleGuideValidationsReportString)
		}
		if strings.Contains(result, consts.ContractSecurityAnalysisReport.Name) {
			var securityAnalysisReport model.ContractCheckResult[[]model.ContractStyleGuideValidationsReportDetails]
			err := json.Unmarshal(file, &securityAnalysisReport)
			if err != nil {
				continue
			}
			var securityAnalysisReportString model.ContractCheckResult[string]
			securityAnalysisReportString.Tool = securityAnalysisReport.Tool
			securityAnalysisReportString.Name = securityAnalysisReport.Name
			securityAnalysisReportString.Result = securityAnalysisReport.Result
			var contractCheckResultDetailsList []model.ContractCheckResultDetails[string]
			for _, report := range securityAnalysisReport.Context {
				var contractCheckResultDetails model.ContractCheckResultDetails[string]
				contractCheckResultDetails.Name = report.Name
				contractCheckResultDetails.Issue = report.Issue
				marshal, err := json.Marshal(report.Message)
				if err != nil {
					continue
				}
				contractCheckResultDetails.Message = string(marshal)
				contractCheckResultDetailsList = append(contractCheckResultDetailsList, contractCheckResultDetails)
			}
			securityAnalysisReportString.Context = contractCheckResultDetailsList
			contractCheckResultList = append(contractCheckResultList, securityAnalysisReportString)
		}
	}
	a.path = path2.Join(destDir, consts.CheckAggregationResult)
	create, err := os.Create(a.path)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(contractCheckResultList)
	if err != nil {
		return nil, err
	}
	_, err = create.WriteString(string(marshal))
	if err != nil {
		return nil, err
	}
	create.Close()

	id, err := strconv.Atoi(jobId)
	if err != nil {
		return nil, err
	}
	actionResult := model.ActionResult{
		Artifactorys: nil,
		Reports: []model.Report{
			{
				Id:   id,
				Url:  a.path,
				Type: 2,
			},
		},
	}
	return &actionResult, err
}

func (a *CheckAggregationAction) Post() error {
	return nil
}
