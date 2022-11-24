package action

import (
	"context"
	"github.com/hamster-shared/a-line/pkg/model"
	"github.com/hamster-shared/a-line/pkg/output"
	"os"
)

type WorkdirAction struct {
	workdir string
	output  *output.Output
	ctx     context.Context
}

func NewWorkdirAction(step model.Step, ctx context.Context, output *output.Output) *WorkdirAction {

	return &WorkdirAction{
		ctx:     ctx,
		output:  output,
		workdir: step.With["workdir"],
	}
}

func (a *WorkdirAction) Pre() error {
	return nil
}

// Hook 执行
func (a *WorkdirAction) Hook() (*model.ActionResult, error) {
	_, err := os.Stat(a.workdir)
	if err != nil {
		a.output.WriteLine(" workdir not exists: " + a.workdir)
		return nil, err
	}

	stack := a.ctx.Value(STACK).(map[string]interface{})
	stack["workdir"] = a.workdir

	return nil, nil
}

// Post 执行后清理 (无论执行是否成功，都应该有Post的清理)
func (a *WorkdirAction) Post() error {
	return nil
}
