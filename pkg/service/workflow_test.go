package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/mohaijiang/agent-go/candid"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"text/template"
	"time"
)

func TestWorkflowService_SyncContract(t *testing.T) {
	path := "/Users/mohaijiang/pipelines/jobs/1_1/artifactory/10/MetaCoin.json"

	data, _ := os.ReadFile(path)

	m := make(map[string]any)

	err := json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m["abi"])
	abi, err := json.Marshal(m["abi"])
	fmt.Println(string(abi))
	fmt.Println(m["bytecode"])
}

func TestTemplate(t *testing.T) {
	filePath := "templates/frontend-deploy.yml"
	content, err := temp.ReadFile(filePath)
	fileContent := string(content)
	tmpl, err := template.New("test").Delims("[[", "]]").Parse(fileContent)
	if err != nil {
		log.Println("template parse failed ", err.Error())
		return
	}
	templateData := parameter.TemplateCheck{
		Name:          "Name",
		RepositoryUrl: "www.baidu.com",
	}
	var input bytes.Buffer
	err = tmpl.Execute(&input, templateData)

	fmt.Println(input.String())
}

func TestGetSuiModelName(t *testing.T) {
	short := strings.TrimPrefix("https://github.com/mohaijiang/my-sui-nft.git", "https://github.com/")
	short = strings.TrimSuffix(short, ".git")
	fmt.Println(short)
	moveUrl := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/Move.toml", short, "main")
	fmt.Println(moveUrl)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get(moveUrl)
	if err != nil {
		panic(err)
		return
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	var config vo.Config
	_, err = toml.Decode(string(data), &config)

	if err != nil {
		panic(err)
	}
	fmt.Println(config.Package.Name)

}

func TestDidToJson(t *testing.T) {
	didUrl := "/Users/abing/GitHub/Hamster/Template/examples/motoko/whoami/.dfx/local/canisters/whoami/whoami.did"
	// analysis did
	didContent, err := readDid(didUrl)
	if err != nil {
		panic(t)
	}

	discription, err := candid.ParseDID([]byte(didContent))
	if err != nil {
		panic(t)
	}

	bytes, err := json.Marshal(discription)
	if err != nil {
		panic(t)
	}
	fmt.Println("-------json----")
	fmt.Println(string(bytes))
	fmt.Println("-------json----")
}
