package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
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
