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

// SolHintAction SolHint合约检查
type SolHintAction struct {
	path   string
	ctx    context.Context
	output *output.Output
}

func NewSolHintAction(step model.Step, ctx context.Context, output *output.Output) *SolHintAction {
	return &SolHintAction{
		path:   step.With["path"],
		ctx:    ctx,
		output: output,
	}
}

func (a *SolHintAction) Pre() error {
	stack := a.ctx.Value(STACK).(map[string]interface{})

	workdir, ok := stack["workdir"].(string)
	if !ok {
		return errors.New("workdir is empty")
	}
	create, err := os.Create(path2.Join(workdir, consts.SolHintCheckInitFileName))
	if err != nil {
		return err
	}
	_, err = create.WriteString(consts.SolHintCheckRule)
	if err != nil {
		return err
	}
	create.Close()
	return nil
}

func (a *SolHintAction) Hook() (*model.ActionResult, error) {

	a.output.NewStep("sol-profiler-check")

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
	absPathList = utils.GetSuffixFiles(path2.Join(workdir, a.path), consts.SolFileSuffix, absPathList)
	destDir := path2.Join(userHomeDir, consts.ArtifactoryDir, jobName, consts.CheckName, jobId, consts.SolHintCheckOutputDir)
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
		command := consts.SolHintCheck + path
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

func (a *SolHintAction) Post() error {
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
	var checkResultDetailsList []model.ContractCheckResultDetails[[]model.ContractStyleGuideValidationsReportDetails]
	successFlag := true
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
			if strings.Contains(s, "  ") {
				split := strings.Split(s, "  ")
				lineAndCol := strings.Split(strings.TrimSpace(split[1]), ":")
				validationsReportDetails := model.NewContractStyleGuideValidationsReportDetails(lineAndCol[0], lineAndCol[1], split[2], "", strings.Join(split[3:len(split)-1], " "), consts.SolHintCheckOutputDir)
				styleGuideValidationsReportDetailsList = append(styleGuideValidationsReportDetailsList, validationsReportDetails)
				successFlag = false
			}
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
	checkResult := model.NewContractCheckResult(consts.ContractStyleGuideValidationsReport.Name, result, consts.ContractStyleGuideValidationsReport.Tool, checkResultDetailsList)
	fmt.Println(checkResult)
	create, err := os.Create(path2.Join(a.path, consts.CheckResult))
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

func (a *SolHintAction) ExecuteCommand(commands []string, workdir string) (string, error) {
	c := exec.CommandContext(a.ctx, commands[0], commands[1:]...) // mac linux
	c.Dir = workdir
	logger.Debugf("execute solhint -f table *.sol command: %s", strings.Join(commands, " "))
	a.output.WriteCommandLine(strings.Join(commands, " "))
	out, err := c.CombinedOutput()
	fmt.Println(string(out))
	a.output.WriteCommandLine(string(out))
	if err != nil {
		a.output.WriteLine(err.Error())
	}
	return string(out), err
}
