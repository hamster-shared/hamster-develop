package action

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/engine/consts"
	"github.com/hamster-shared/a-line/engine/logger"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/engine/output"
	"github.com/hamster-shared/a-line/pkg/utils"
	"io"
	"os"
	"os/exec"
	path2 "path"
	"strconv"
	"strings"
)

// MythRilAction mythRil合约检查
type MythRilAction struct {
	path        string
	solcVersion string
	ctx         context.Context
	output      *output.Output
}

func NewMythRilAction(step model.Step, ctx context.Context, output *output.Output) *MythRilAction {
	return &MythRilAction{
		path:        step.With["path"],
		solcVersion: step.With["solc-version"],
		ctx:         ctx,
		output:      output,
	}
}

func (a *MythRilAction) Pre() error {
	return nil
}

func (a *MythRilAction) Hook() (*model.ActionResult, error) {

	a.output.NewStep("mythril-check")

	stack := a.ctx.Value(STACK).(map[string]interface{})

	workdir, ok := stack["workdir"].(string)
	if !ok {
		return nil, errors.New("workdir is empty")
	}
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
	basePath := path2.Join(workdir, a.path)
	absPathList = utils.GetSuffixFiles(basePath, consts.SolFileSuffix, absPathList)
	destDir := path2.Join(userHomeDir, consts.ArtifactoryDir, jobName, consts.CheckName, jobId, consts.MythRilCheckOutputDir)
	_, err = os.Stat(destDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	for _, path := range absPathList {
		_, filenameOnly := utils.GetFilenameWithSuffixAndFilenameOnly(path)
		dest := path2.Join(destDir, filenameOnly+consts.SuffixType)
		err, redundantPath := utils.GetRedundantPath(basePath, path)
		if err != nil {
			return nil, err
		}
		commandTemplate := consts.MythRilCheck
		command := fmt.Sprintf(commandTemplate, basePath, redundantPath)
		if a.solcVersion != "" {
			command = command + " --solv " + a.solcVersion
		}
		fields := strings.Fields(command)
		out, err := a.ExecuteCommand(fields, workdir)
		if err != nil {
			return nil, err
		}
		create, err := os.Create(dest)
		if err != nil {
			return nil, err
		}
		_, err = create.WriteString(out)
		if err != nil {
			return nil, err
		}
		create.Close()
	}
	a.path = destDir
	id, err := strconv.Atoi(jobId)
	if err != nil {
		return nil, err
	}
	actionResult := model.ActionResult{
		Artifactorys: nil,
		Reports: []model.Report{
			{
				Id:   id,
				Url:  "",
				Type: 2,
			},
		},
	}
	return &actionResult, err
}

func (a *MythRilAction) Post() error {
	open, err := os.Open(a.path)
	if err != nil {
		return err
	}
	fileInfo, err := open.Stat()
	if err != nil {
		return err
	}
	isDir := fileInfo.IsDir()
	if !isDir {
		return errors.New("check result path is err")
	}
	fileInfos, err := open.Readdir(-1)
	successFlag := true
	var checkResultDetailsList []model.ContractCheckResultDetails[[]model.ContractStyleGuideValidationsReportDetails]
	for _, info := range fileInfos {
		path := path2.Join(a.path, info.Name())
		var styleGuideValidationsReportDetailsList []model.ContractStyleGuideValidationsReportDetails
		file, err := os.Open(path)
		if err != nil {
			return errors.New("file open fail")
		}
		defer file.Close()

		line := bufio.NewReader(file)
		for {
			content, _, err := line.ReadLine()
			if err == io.EOF {
				break
			}
			s := string(content)
			if strings.Contains(s, "The analysis was completed successfully") {
				break
			}
			var styleGuideValidationsReportDetails model.ContractStyleGuideValidationsReportDetails
			styleGuideValidationsReportDetails.Tool = consts.MythRilCheckOutputDir
			if strings.Contains(s, "Error:") || strings.Contains(s, "Note:") {
				index := strings.Index(s, ":")
				styleGuideValidationsReportDetails.Note = s[index+2:]
			}
			if strings.Contains(s, "-->") {
				split := strings.Split(s, ":")
				if len(split) < 3 {
					continue
				}
				styleGuideValidationsReportDetails.Line = split[1]
				styleGuideValidationsReportDetails.Column = split[2]
			}
			if styleGuideValidationsReportDetails.Note == "" && styleGuideValidationsReportDetails.Line == "" {
				continue
			}
			if styleGuideValidationsReportDetails.Note == "" && len(styleGuideValidationsReportDetailsList) >= 1 {
				styleGuideValidationsReportDetailsList[len(styleGuideValidationsReportDetailsList)-1].Line = styleGuideValidationsReportDetails.Line
				styleGuideValidationsReportDetailsList[len(styleGuideValidationsReportDetailsList)-1].Column = styleGuideValidationsReportDetails.Column
			} else {
				styleGuideValidationsReportDetailsList = append(styleGuideValidationsReportDetailsList, styleGuideValidationsReportDetails)
			}
			successFlag = false
		}
		details := model.NewContractCheckResultDetails[[]model.ContractStyleGuideValidationsReportDetails](strings.Replace(info.Name(), consts.SuffixType, consts.SolFileSuffix, 1), len(styleGuideValidationsReportDetailsList), styleGuideValidationsReportDetailsList)
		checkResultDetailsList = append(checkResultDetailsList, details)
	}
	var result string
	if successFlag {
		result = consts.CheckSuccess.Result
	} else {
		result = consts.CheckFail.Result
	}
	checkResult := model.NewContractCheckResult(consts.ContractSecurityAnalysisReport.Name, result, consts.ContractSecurityAnalysisReport.Tool, checkResultDetailsList)
	create, err := os.Create(path2.Join(a.path, consts.CheckResult))
	fmt.Println(checkResult)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(checkResult)
	if err != nil {
		return err
	}
	_, err = create.WriteString(string(marshal))
	if err != nil {
		return err
	}
	create.Close()
	return nil
}

func (a *MythRilAction) ExecuteCommand(commands []string, workdir string) (string, error) {
	c := exec.CommandContext(a.ctx, commands[0], commands[1:]...) // mac linux
	c.Dir = workdir
	logger.Debugf("execute mythril check command: %s", strings.Join(commands, " "))
	a.output.WriteCommandLine(strings.Join(commands, " "))
	out, err := c.CombinedOutput()
	fmt.Println(string(out))
	a.output.WriteCommandLine(string(out))
	if err != nil {
		a.output.WriteLine(err.Error())
	}
	return string(out), err
}
