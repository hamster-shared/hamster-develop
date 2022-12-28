package action

import (
	"context"
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
	return nil, err
}

func (a *SolProfilerAction) Post() error {
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
