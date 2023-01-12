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
	"regexp"
	"strconv"
	"strings"
)

// SolProfilerAction SolProfiler合约检查
type SolProfilerAction struct {
	path   string
	ctx    context.Context
	output *output.Output
}

func NewSolProfilerAction(step model.Step, ctx context.Context, output *output.Output) *SolProfilerAction {
	return &SolProfilerAction{
		path:   step.With["path"],
		ctx:    ctx,
		output: output,
	}
}

func (a *SolProfilerAction) Pre() error {
	return nil
}

func (a *SolProfilerAction) Hook() (*model.ActionResult, error) {

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
	destDir := path2.Join(userHomeDir, consts.ArtifactoryDir, jobName, consts.CheckName, jobId, consts.SolProfilerCheckOutputDir)
	_, err = os.Stat(destDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	for _, path := range absPathList {
		_, file := path2.Split(path)
		var filenameWithSuffix string
		filenameWithSuffix = path2.Base(file)
		var fileSuffix string
		fileSuffix = path2.Ext(filenameWithSuffix)
		var filenameOnly string
		filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)

		dest := path2.Join(destDir, filenameOnly+consts.SuffixType)
		command := consts.SolProfilerCheck + path
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

func (a *SolProfilerAction) Post() error {
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
	var checkResultDetailsList []model.ContractCheckResultDetails[[]model.ContractMethodsPropertiesReportDetails]
	compile, err := regexp.Compile(`\[.{1,2}m`)
	if err != nil {
		logger.Errorf("regexp err %s", err)
		return err
	}
	for _, info := range fileInfos {
		path := path2.Join(a.path, info.Name())
		var methodsPropertiesReportDetailsList []model.ContractMethodsPropertiesReportDetails
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
			contentSting := string(content)
			s := compile.ReplaceAllString(contentSting, "")
			if strings.Contains(s, "┌─") || strings.Contains(s, "├─") || strings.Contains(s, "└─") || strings.Contains(s, "Contract/Library/Interface") {
				continue
			}
			if !strings.Contains(s, "│") {
				continue
			}
			split := strings.Split(s, "│")
			trimSpace := strings.TrimSpace(split[1])
			if trimSpace == "" {
				continue
			}
			var methodsPropertiesReportDetails model.ContractMethodsPropertiesReportDetails
			if strings.Contains(trimSpace, "(library)") {
				index := strings.Index(trimSpace, "(library)")
				methodsPropertiesReportDetails.Contract = strings.TrimSpace(trimSpace[:index])
				methodsPropertiesReportDetails.Category = "library"
			} else {
				methodsPropertiesReportDetails.Contract = trimSpace
				methodsPropertiesReportDetails.Category = "Contract"
			}
			methodsPropertiesReportDetails.Function = strings.TrimSpace(split[2])
			methodsPropertiesReportDetails.Visibility = strings.TrimSpace(split[3])
			methodsPropertiesReportDetails.Visibility = strings.TrimSpace(split[4])
			methodsPropertiesReportDetails.Returns = strings.TrimSpace(split[5])
			methodsPropertiesReportDetails.Modifiers = strings.TrimSpace(split[6])
			methodsPropertiesReportDetailsList = append(methodsPropertiesReportDetailsList, methodsPropertiesReportDetails)
		}

		details := model.NewContractCheckResultDetails[[]model.ContractMethodsPropertiesReportDetails](strings.Replace(info.Name(), consts.SuffixType, consts.SolFileSuffix, 1), 0, methodsPropertiesReportDetailsList)
		checkResultDetailsList = append(checkResultDetailsList, details)
	}
	checkResult := model.NewContractCheckResult(consts.ContractMethodsPropertiesReport.Name, consts.CheckSuccess.Result, consts.ContractMethodsPropertiesReport.Tool, checkResultDetailsList)
	fmt.Printf("%+v", checkResult)
	create, err := os.Create(path2.Join(a.path, consts.CheckResult))
	fmt.Println()
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

func (a *SolProfilerAction) ExecuteCommand(commands []string, workdir string) (string, error) {
	c := exec.CommandContext(a.ctx, commands[0], commands[1:]...) // mac linux
	c.Dir = workdir
	logger.Debugf("execute sol-profiler *.sol command: %s", strings.Join(commands, " "))
	a.output.WriteCommandLine(strings.Join(commands, " "))
	out, err := c.CombinedOutput()
	fmt.Println(string(out))
	a.output.WriteCommandLine(string(out))
	if err != nil {
		a.output.WriteLine(err.Error())
	}
	return string(out), err
}
