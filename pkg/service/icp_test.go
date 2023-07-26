package service

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestIcpService_CreateIdentity(t *testing.T) {
	useIdentityCmd := "dfx identity use test1113"
	_, err := execDfxCommand(useIdentityCmd)
	if err != nil {
		panic(err)
	}
	accountIdCmd := "dfx ledger account-id"
	accountId, err := execDfxCommand(accountIdCmd)
	if err != nil {
		panic(err)
	}
	pIdCmd := "dfx identity get-principal"
	pId, err := execDfxCommand(pIdCmd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("identity: %s \naccount-id: %s \nprincipal-id: %s \n", "test1113", strings.TrimSpace(accountId), strings.TrimSpace(pId))
}

func execDfxCommand(cmd string) (string, error) {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if exitError, ok := err.(*exec.ExitError); ok {
		fmt.Errorf("%s Exit status: %d, Exit str: %s", cmd, exitError.ExitCode(), string(exitError.Stderr))
		return "", err
	} else if err != nil {
		// 输出其他类型的错误
		fmt.Errorf("%s Failed to execute command: %s", cmd, err)
		return "", err
	}
	return string(output), nil
}
