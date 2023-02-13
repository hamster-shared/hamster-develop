package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"log"
	"os"
	"testing"
	"text/template"
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
