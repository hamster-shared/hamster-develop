package controller

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
	"log"
	"strconv"
	"time"
)

func (h *HandlerServer) workflowList(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	log.Println(id)
	typeStr := gin.DefaultQuery("type", "0")
	workflowType, err := strconv.Atoi(typeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	log.Println(workflowType)
	var data vo.WorkflowPage
	var workflowData []vo.WorkflowVo
	data1 := vo.WorkflowVo{
		Id:          1,
		ProjectId:   1,
		Type:        1,
		ExecNumber:  1,
		StageInfo:   "id: 0\njob:\n    version: \\\"1\\\"\n    name: my-erc201\n    stages:\n        compile:\n            steps:\n                - name: compile\n                  run: |\n                    npm install\n                    truffle compile\n            needs:\n                - template-init\n        compile-node:\n            steps:\n                - name: compile\n                  run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n            needs:\n                - deploy-contract\n        contract-test:\n            steps:\n                - name: deploy\n                  run: |\n                    truffle test\n            needs:\n                - ganache\n        deploy-contract:\n            steps:\n                - name: deploy-contract\n                  run: |\n                    truffle deploy\n            needs:\n                - ganache\n        deploy-frontend:\n            steps:\n                - name: deploy\n                  run: |\n                    cd app\n                    if [ -f \\\"node.pid\\\" ]; then\n                      kill -9 `cat node.pid`  || (echo 'No such process ')\n                    fi\n                    nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                    sleep 2\n            needs:\n                - compile-node\n        ganache:\n            steps:\n                - name: ganache\n                  run: |\n                    npm install -g ganache\n                    if [ -f \\\"command.pid\\\" ]; then\n                      kill -9 `cat command.pid`  || (echo 'No such process ')\n                    fi\n                    nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                    sleep 2\n            needs:\n                - compile\n        solidity-lint:\n            steps:\n                - name: solidity-check\n                  run: |\n                    npm install -g ethlint\n                    solium --init\n                    solium -d contracts/\n            needs:\n                - compile\n        template-init:\n            steps:\n                - name: set workdir\n                  uses: workdir\n                  with:\n                    workdir: /Users/sunjianguo/tmp/my-erc20\n                - name: template init\n                  uses: git-checkout\n                  with:\n                    branch: main\n                    url: https://github.com/jian-guo-s/truffle-webpack.git\nstatus: 2\ntriggerMode: Manual trigger\nstages:\n    - name: template-init\n      stage:\n        steps:\n            - name: set workdir\n              uses: workdir\n              with:\n                workdir: /Users/sunjianguo/tmp/my-erc20\n            - name: template init\n              uses: git-checkout\n              with:\n                branch: main\n                url: https://github.com/jian-guo-s/truffle-webpack.git\n      status: 3\n      starttime: 2022-12-01T17:34:44.410371+08:00\n      duration: 1330\n    - name: compile\n      stage:\n        steps:\n            - name: compile\n              run: |\n                npm install\n                truffle compile\n        needs:\n            - template-init\n      status: 2\n      starttime: 2022-12-01T17:34:45.741188+08:00\n      duration: 386\n    - name: solidity-lint\n      stage:\n        steps:\n            - name: solidity-check\n              run: |\n                npm install -g ethlint\n                solium --init\n                solium -d contracts/\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: ganache\n      stage:\n        steps:\n            - name: ganache\n              run: |\n                npm install -g ganache\n                if [ -f \\\"command.pid\\\" ]; then\n                  kill -9 `cat command.pid`  || (echo 'No such process ')\n                fi\n                nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                sleep 2\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: contract-test\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                truffle test\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-contract\n      stage:\n        steps:\n            - name: deploy-contract\n              run: |\n                truffle deploy\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: compile-node\n      stage:\n        steps:\n            - name: compile\n              run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n        needs:\n            - deploy-contract\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-frontend\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                cd app\n                if [ -f \\\"node.pid\\\" ]; then\n                  kill -9 `cat node.pid`  || (echo 'No such process ')\n                fi\n                nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                sleep 2\n        needs:\n            - compile-node\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\nstartTime: 2022-12-01T17:34:44.409461+08:00\nduration: 1723\nactionResult:\n    artifactorys: []\n    reports: []\noutput:\n    name: my-erc201\n    id: 0",
		TriggerUser: "张三",
		CodeInfo:    "main | b4a0d99",
		TriggerMode: 1,
		Status:      2,
		StartTime:   time.Now().AddDate(0, 0, -1),
	}
	data2 := vo.WorkflowVo{
		Id:          2,
		ProjectId:   2,
		Type:        2,
		ExecNumber:  1,
		StageInfo:   "id: 0\njob:\n    version: \\\"1\\\"\n    name: my-erc201\n    stages:\n        compile:\n            steps:\n                - name: compile\n                  run: |\n                    npm install\n                    truffle compile\n            needs:\n                - template-init\n        compile-node:\n            steps:\n                - name: compile\n                  run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n            needs:\n                - deploy-contract\n        contract-test:\n            steps:\n                - name: deploy\n                  run: |\n                    truffle test\n            needs:\n                - ganache\n        deploy-contract:\n            steps:\n                - name: deploy-contract\n                  run: |\n                    truffle deploy\n            needs:\n                - ganache\n        deploy-frontend:\n            steps:\n                - name: deploy\n                  run: |\n                    cd app\n                    if [ -f \\\"node.pid\\\" ]; then\n                      kill -9 `cat node.pid`  || (echo 'No such process ')\n                    fi\n                    nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                    sleep 2\n            needs:\n                - compile-node\n        ganache:\n            steps:\n                - name: ganache\n                  run: |\n                    npm install -g ganache\n                    if [ -f \\\"command.pid\\\" ]; then\n                      kill -9 `cat command.pid`  || (echo 'No such process ')\n                    fi\n                    nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                    sleep 2\n            needs:\n                - compile\n        solidity-lint:\n            steps:\n                - name: solidity-check\n                  run: |\n                    npm install -g ethlint\n                    solium --init\n                    solium -d contracts/\n            needs:\n                - compile\n        template-init:\n            steps:\n                - name: set workdir\n                  uses: workdir\n                  with:\n                    workdir: /Users/sunjianguo/tmp/my-erc20\n                - name: template init\n                  uses: git-checkout\n                  with:\n                    branch: main\n                    url: https://github.com/jian-guo-s/truffle-webpack.git\nstatus: 2\ntriggerMode: Manual trigger\nstages:\n    - name: template-init\n      stage:\n        steps:\n            - name: set workdir\n              uses: workdir\n              with:\n                workdir: /Users/sunjianguo/tmp/my-erc20\n            - name: template init\n              uses: git-checkout\n              with:\n                branch: main\n                url: https://github.com/jian-guo-s/truffle-webpack.git\n      status: 3\n      starttime: 2022-12-01T17:34:44.410371+08:00\n      duration: 1330\n    - name: compile\n      stage:\n        steps:\n            - name: compile\n              run: |\n                npm install\n                truffle compile\n        needs:\n            - template-init\n      status: 2\n      starttime: 2022-12-01T17:34:45.741188+08:00\n      duration: 386\n    - name: solidity-lint\n      stage:\n        steps:\n            - name: solidity-check\n              run: |\n                npm install -g ethlint\n                solium --init\n                solium -d contracts/\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: ganache\n      stage:\n        steps:\n            - name: ganache\n              run: |\n                npm install -g ganache\n                if [ -f \\\"command.pid\\\" ]; then\n                  kill -9 `cat command.pid`  || (echo 'No such process ')\n                fi\n                nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                sleep 2\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: contract-test\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                truffle test\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-contract\n      stage:\n        steps:\n            - name: deploy-contract\n              run: |\n                truffle deploy\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: compile-node\n      stage:\n        steps:\n            - name: compile\n              run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n        needs:\n            - deploy-contract\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-frontend\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                cd app\n                if [ -f \\\"node.pid\\\" ]; then\n                  kill -9 `cat node.pid`  || (echo 'No such process ')\n                fi\n                nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                sleep 2\n        needs:\n            - compile-node\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\nstartTime: 2022-12-01T17:34:44.409461+08:00\nduration: 1723\nactionResult:\n    artifactorys: []\n    reports: []\noutput:\n    name: my-erc201\n    id: 0",
		TriggerUser: "李四",
		CodeInfo:    "main | b4a0d99",
		TriggerMode: 1,
		Status:      1,
		StartTime:   time.Now().AddDate(0, 0, -2),
	}
	workflowData = append(workflowData, data1, data2)
	data.Data = workflowData
	data.Page = 1
	data.PageSize = 10
	Success(workflowData, gin)
}

func (h *HandlerServer) workflowDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	log.Println(workflowId)
	data := vo.WorkflowDetailVo{
		Id:         1,
		WorkflowId: 1,
		StageInfo:  "id: 0\njob:\n    version: \\\"1\\\"\n    name: my-erc201\n    stages:\n        compile:\n            steps:\n                - name: compile\n                  run: |\n                    npm install\n                    truffle compile\n            needs:\n                - template-init\n        compile-node:\n            steps:\n                - name: compile\n                  run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n            needs:\n                - deploy-contract\n        contract-test:\n            steps:\n                - name: deploy\n                  run: |\n                    truffle test\n            needs:\n                - ganache\n        deploy-contract:\n            steps:\n                - name: deploy-contract\n                  run: |\n                    truffle deploy\n            needs:\n                - ganache\n        deploy-frontend:\n            steps:\n                - name: deploy\n                  run: |\n                    cd app\n                    if [ -f \\\"node.pid\\\" ]; then\n                      kill -9 `cat node.pid`  || (echo 'No such process ')\n                    fi\n                    nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                    sleep 2\n            needs:\n                - compile-node\n        ganache:\n            steps:\n                - name: ganache\n                  run: |\n                    npm install -g ganache\n                    if [ -f \\\"command.pid\\\" ]; then\n                      kill -9 `cat command.pid`  || (echo 'No such process ')\n                    fi\n                    nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                    sleep 2\n            needs:\n                - compile\n        solidity-lint:\n            steps:\n                - name: solidity-check\n                  run: |\n                    npm install -g ethlint\n                    solium --init\n                    solium -d contracts/\n            needs:\n                - compile\n        template-init:\n            steps:\n                - name: set workdir\n                  uses: workdir\n                  with:\n                    workdir: /Users/sunjianguo/tmp/my-erc20\n                - name: template init\n                  uses: git-checkout\n                  with:\n                    branch: main\n                    url: https://github.com/jian-guo-s/truffle-webpack.git\nstatus: 2\ntriggerMode: Manual trigger\nstages:\n    - name: template-init\n      stage:\n        steps:\n            - name: set workdir\n              uses: workdir\n              with:\n                workdir: /Users/sunjianguo/tmp/my-erc20\n            - name: template init\n              uses: git-checkout\n              with:\n                branch: main\n                url: https://github.com/jian-guo-s/truffle-webpack.git\n      status: 3\n      starttime: 2022-12-01T17:34:44.410371+08:00\n      duration: 1330\n    - name: compile\n      stage:\n        steps:\n            - name: compile\n              run: |\n                npm install\n                truffle compile\n        needs:\n            - template-init\n      status: 2\n      starttime: 2022-12-01T17:34:45.741188+08:00\n      duration: 386\n    - name: solidity-lint\n      stage:\n        steps:\n            - name: solidity-check\n              run: |\n                npm install -g ethlint\n                solium --init\n                solium -d contracts/\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: ganache\n      stage:\n        steps:\n            - name: ganache\n              run: |\n                npm install -g ganache\n                if [ -f \\\"command.pid\\\" ]; then\n                  kill -9 `cat command.pid`  || (echo 'No such process ')\n                fi\n                nohup ganache > ganache.log 2>&1& echo $! > command.pid\n                sleep 2\n        needs:\n            - compile\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: contract-test\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                truffle test\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-contract\n      stage:\n        steps:\n            - name: deploy-contract\n              run: |\n                truffle deploy\n        needs:\n            - ganache\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: compile-node\n      stage:\n        steps:\n            - name: compile\n              run: \\\"cd app && npm install\\\\nnpm run build          \\\\n\\\"\n        needs:\n            - deploy-contract\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\n    - name: deploy-frontend\n      stage:\n        steps:\n            - name: deploy\n              run: |\n                cd app\n                if [ -f \\\"node.pid\\\" ]; then\n                  kill -9 `cat node.pid`  || (echo 'No such process ')\n                fi\n                nohup  npm run dev  > node.log 2>&1& echo $! > node.pid\n                sleep 2\n        needs:\n            - compile-node\n      status: 0\n      starttime: 0001-01-01T00:00:00Z\n      duration: 0\nstartTime: 2022-12-01T17:34:44.409461+08:00\nduration: 1723\nactionResult:\n    artifactorys: []\n    reports: []\noutput:\n    name: my-erc201\n    id: 0",
		Status:     1,
		StartTime:  time.Now().AddDate(0, 0, -1),
		EndTime:    time.Now(),
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowContract(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	log.Println(workflowDetailId)
	log.Println(workflowId)
	var data []db2.Contract
	data1 := db2.Contract{
		Id:               1,
		WorkflowId:       1,
		WorkflowDetailId: 1,
		Name:             "contract-one",
		Version:          "#1",
		Network:          "mainnet",
		BuildTime:        time.Now().AddDate(0, 0, -1),
	}
	data2 := db2.Contract{
		Id:               1,
		WorkflowId:       1,
		WorkflowDetailId: 1,
		Name:             "contract-two",
		Version:          "#1",
		Network:          "Testnet/Goerli",
		BuildTime:        time.Now().AddDate(0, 0, -1),
	}
	data = append(data, data1, data2)
	Success(data, gin)
}

func (h *HandlerServer) workflowReport(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	log.Println(workflowDetailId)
	log.Println(workflowId)
	var data []vo.ReportVo
	data1 := vo.ReportVo{
		Id:               1,
		WorkflowId:       1,
		WorkflowDetailId: 2,
		ProjectId:        1,
		Name:             "report-one",
		Type:             1,
		CheckTool:        "Ethlint",
		Result:           "2 errors found",
		CheckTime:        time.Now().AddDate(0, 0, -1),
		ReportFile:       "",
	}
	data2 := vo.ReportVo{
		Id:               2,
		WorkflowId:       1,
		WorkflowDetailId: 2,
		ProjectId:        1,
		Name:             "report-two",
		Type:             1,
		CheckTool:        "Ethlint",
		Result:           "2 errors found",
		CheckTime:        time.Now().AddDate(0, 0, -1),
		ReportFile:       "",
	}
	data = append(data, data1, data2)
	Success(data, gin)
}
