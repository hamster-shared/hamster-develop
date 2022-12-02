package action

import (
	"context"
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/model"
	"github.com/hamster-shared/a-line/pkg/output"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
	"os"
)

// TruffleDeployAction truffle deploy action
type TruffleDeployAction struct {
	network    string // deploy network
	privateKey string // deploy Private key
	output     *output.Output
	ctx        context.Context
}

func NewTruffleDeployAction(step model.Step, ctx context.Context, output *output.Output) *TruffleDeployAction {
	return &TruffleDeployAction{
		network:    step.With["network"],
		privateKey: step.With["private-key"],
		ctx:        ctx,
		output:     output,
	}
}

func (a *TruffleDeployAction) Pre() error {
	stack := a.ctx.Value(STACK).(map[string]interface{})
	workdir := stack["workdir"].(string)
	_, err := os.Stat(workdir)
	if os.IsNotExist(err) {
		return errors.New("workdir not exist")
	}
	if a.network != "default" {
		if a.privateKey == "" {
			return errors.New("workdir not exist")
		}
	}
	log.Println("---------------")
	log.Println(workdir)
	log.Println("---------------")
	newShell := shell.NewShell(consts.IPFS_SHELL)
	version, s, err := newShell.Version()
	if err != nil {
		return errors.New("get workdir error")
	}
	fmt.Println(fmt.Sprintf("ipfs version is %s, commit sha is %s", version, s))
	return nil
}
