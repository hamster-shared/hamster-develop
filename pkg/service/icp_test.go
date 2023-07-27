package service

import (
	"fmt"
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid"
	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/ledger"
	"github.com/aviate-labs/agent-go/principal"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
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

func TestIcpService_InitWallet(t *testing.T) {
	str := `
Transfer sent at BlockHeight: 20
Canister created with id: "gastn-uqaaa-aaaae-aaafq-cai"
`

	// 定义正则表达式
	re := regexp.MustCompile(`Canister created with id: "(.*?)"`)

	// 使用正则表达式提取字符串
	matches := re.FindStringSubmatch(str)

	// 获取匹配到的字符串
	if len(matches) > 1 {
		value := matches[1]
		fmt.Println("Value:", value)
	} else {
		fmt.Println("String not found.")
	}
}

var ic0, _ = url.Parse("https://ic0.app/")

func TestIcpService_AgentGo(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("failed to get user home directory %s", err)
		panic(t)
	}

	// 构建文件路径
	filePath := filepath.Join(homeDir, ".config", "dfx", "identity", "test1115", "identity.pem")

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("failed to read file %s", err)
		panic(t)
	}
	fmt.Printf("dfx 生成的pem是\n  %s \n", string(content))
	//从pem文件获取新的标识
	ed25519Identity, _ := identity.NewRandomEd25519Identity()
	pem, _ := ed25519Identity.ToPEM()
	fmt.Printf("sdk 生成的pem是\n  %s \n", string(pem))
	id, err := identity.NewEd25519IdentityFromPEM(pem)

	//获取控制者
	c := agent.NewClient(agent.ClientConfig{Host: ic0})
	status, _ := c.Status()
	fmt.Println(status.Version)
	canisterId := "ryjl3-tyaaa-aaaaa-aaaba-cai"
	ledgerID, _ := principal.Decode(canisterId)
	a := agent.New(agent.Config{
		Identity:     id,
		ClientConfig: &agent.ClientConfig{Host: ic0},
	})

	//获取控制者
	controllers, err := a.GetCanisterControllers(ledgerID)
	if err != nil {
		fmt.Println(err)
		panic(t)
	}
	fmt.Printf("%s 的控制者是 %s \n", canisterId, controllers)

	//查询方法
	args, err := candid.EncodeValueString("record { account = \"9523dc824aa062dcd9c91b98f4594ff9c6af661ac96747daef2090b7fe87037d\" }")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.QueryString(ledgerID, "account_balance_dfx", args))

}
func TestIcpService_AgentGo_Fail(t *testing.T) {
	canisterId := "aaaaa-aa"
	ledgerID, _ := principal.Decode(canisterId)
	ed25519Identity, _ := identity.NewRandomEd25519Identity()
	a := agent.New(agent.Config{
		Identity:     ed25519Identity,
		ClientConfig: &agent.ClientConfig{Host: ic0},
	})

	//查询方法
	args, err := candid.EncodeValueString("record { canister_id = \"ryjl3-tyaaa-aaaaa-aaaba-cai\" }")
	if err != nil {
		fmt.Println(err)
		panic(t)
	}
	callRaw, err := a.QueryString(ledgerID, "canister_status", args)
	if err != nil {
		fmt.Println(err)
		panic(t)
	}
	fmt.Println(callRaw)
}

func TestIcpService_AgentGo_Transfer(t *testing.T) {
	canisterId := "ryjl3-tyaaa-aaaaa-aaaba-cai"
	ledgerID, _ := principal.Decode(canisterId)
	a := ledger.New(ledgerID, ic0)
	p, _ := principal.Decode("aaaaa-aa")
	subAccount := ledger.SubAccount(principal.DefaultSubAccount)
	tokens, err := a.Transfer(ledger.TransferArgs{
		Memo: 0,
		Amount: ledger.Tokens{
			E8S: 100_000,
		},
		Fee: ledger.Tokens{
			E8S: 10_000,
		},
		FromSubAccount: &subAccount,
		To:             p.AccountIdentifier(principal.DefaultSubAccount),
		CreatedAtTime: &ledger.TimeStamp{
			TimestampNanos: uint64(time.Now().UnixNano()),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if *tokens != 1 {
		t.Error(tokens)
	}
}
