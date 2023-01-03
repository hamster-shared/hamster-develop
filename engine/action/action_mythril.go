package action

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/engine/consts"
	"github.com/hamster-shared/a-line/engine/logger"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/engine/output"
	"github.com/hamster-shared/a-line/pkg/utils"
	"os"
	"os/exec"
	path2 "path"
	"strings"
)

// MythRilAction mythRil合约检查
type MythRilAction struct {
	path   string
	ctx    context.Context
	output *output.Output
}

func NewMythRilAction(step model.Step, ctx context.Context, output *output.Output) *MythRilAction {
	return &MythRilAction{
		path:   step.With["path"],
		ctx:    ctx,
		output: output,
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
	return nil, err
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
	var checkResultDetailsList []model.ContractCheckResultDetails
	for _, info := range fileInfos {
		path := path2.Join(a.path, info.Name())
		file, err := os.ReadFile(path)
		if err != nil {
			return errors.New("file open fail")
		}
		result := string(file)
		if !strings.Contains(result, "The analysis was completed successfully") {
			successFlag = false
		}
		details := model.NewContractCheckResultDetails(strings.Replace(info.Name(), consts.SuffixType, consts.SolFileSuffix, 1), result)
		checkResultDetailsList = append(checkResultDetailsList, details)
	}
	var result string
	if successFlag {
		result = consts.CheckSuccess.Result
	} else {
		result = consts.CheckFail.Result
	}
	checkResult := model.NewContractCheckResult(consts.MythRil.Name, result, consts.MythRil.Tool, checkResultDetailsList)
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
