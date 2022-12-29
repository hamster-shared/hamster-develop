package engine

import (
	"fmt"
	"github.com/hamster-shared/a-line/engine/dispatcher"
	"github.com/hamster-shared/a-line/engine/executor"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/engine/service"
	"os"
)

type Engine struct {
	channel         chan model.QueueMessage
	callbackChannel chan model.StatusChangeMessage
	jobService      service.IJobService
	dispatch        dispatcher.IDispatcher
	executeClient   *executor.ExecutorClient
}

func NewEngine() *Engine {

	channel := make(chan model.QueueMessage)
	callbackChannel := make(chan model.StatusChangeMessage)
	jobService := service.NewJobService()
	dispatch := dispatcher.NewDispatcher(channel, callbackChannel)
	executeClient := executor.NewExecutorClient(channel, callbackChannel, jobService)

	hostname, _ := os.Hostname()

	dispatch.Register(&model.Node{
		Name:    hostname,
		Address: "127.0.0.1",
	})

	return &Engine{
		channel:         channel,
		callbackChannel: callbackChannel,
		jobService:      jobService,
		dispatch:        dispatch,
		executeClient:   executeClient,
	}
}

func (e *Engine) Start() {
	e.executeClient.Main()
}

func (e *Engine) CreateJob(name string, yaml string) error {

	return e.jobService.SaveJob(name, yaml)
}

func (e *Engine) DeleteJob(name string) error {
	return e.jobService.DeleteJob(name)
}

func (e *Engine) UpdateJob(name, newName, jobYaml string) error {

	return e.jobService.UpdateJob(name, newName, jobYaml)
}

func (e *Engine) GetJob(name string) *model.Job {
	return e.jobService.GetJobObject(name)
}

func (e *Engine) GetJobs(keyword string, page int, size int) *model.JobPage {
	return e.jobService.JobList(keyword, page, size)
}

func (e *Engine) ExecuteJob(name string) (*model.JobDetail, error) {

	job := e.jobService.GetJobObject(name)
	jobDetail, err := e.jobService.ExecuteJob(name)
	if err != nil {
		return nil, err
	}
	node := e.dispatch.DispatchNode(job)
	e.dispatch.SendJob(jobDetail, node)
	return jobDetail, nil
}

func (e *Engine) ReExecuteJob(name string, historyId int) error {

	err := e.jobService.ReExecuteJob(name, historyId)
	job := e.jobService.GetJobObject(name)
	jobDetail := e.jobService.GetJobDetail(name, historyId)
	node := e.dispatch.DispatchNode(job)
	e.dispatch.SendJob(jobDetail, node)
	return err
}

func (e *Engine) TerminalJob(name string, historyId int) error {

	err := e.jobService.StopJobDetail(name, historyId)
	if err != nil {
		return err
	}
	job := e.jobService.GetJobObject(name)
	jobDetail := e.jobService.GetJobDetail(name, historyId)
	node := e.dispatch.DispatchNode(job)
	e.dispatch.CancelJob(jobDetail, node)
	return nil
}

func (e *Engine) GetJobHistory(name string, historyId int) *model.JobDetail {
	return e.jobService.GetJobDetail(name, historyId)
}

func (e *Engine) GetJobHistorys(name string, page, size int) *model.JobDetailPage {
	return e.jobService.JobDetailList(name, page, size)
}

func (e *Engine) DeleteJobHistory(name string, historyId int) error {
	return e.jobService.DeleteJobDetail(name, historyId)
}

func (e *Engine) GetJobHistoryLog(name string, historyId int) *model.JobLog {
	return e.jobService.GetJobLog(name, historyId)
}

func (e *Engine) GetJobHistoryStageLog(name string, historyId int, stageName string, start int) *model.JobStageLog {
	return e.jobService.GetJobStageLog(name, historyId, stageName, start)
}

func (e *Engine) GetCodeInfo(name string, historyId int) string {
	jobDetail := e.jobService.GetJobDetail(name, historyId)
	if jobDetail != nil {
		return jobDetail.CodeInfo
	}
	return ""
}

func (e *Engine) RegisterStatusChangeHook(hookResult func(message model.StatusChangeMessage)) {
	for { //

		//3. 监听队列
		statusMsg, ok := <-e.callbackChannel
		if !ok {
			return
		}

		fmt.Println("=======[status callback]=========")
		fmt.Println(statusMsg)
		fmt.Println("=======[status callback]=========")

		if hookResult != nil {
			hookResult(statusMsg)
		}
	}
}
