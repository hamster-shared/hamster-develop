package model

type Command int

const (
	Command_Start Command = iota
	Command_Stop
)

type QueueMessage struct {
	ProjectName string
	JobName     string
	JobId       int
	Command     Command
}

func NewStartQueueMsg(projectName, name string, id int) QueueMessage {
	return QueueMessage{
		ProjectName: projectName,
		JobName:     name,
		JobId:       id,
		Command:     Command_Start,
	}

}

func NewStopQueueMsg(projectName, name string, id int) QueueMessage {
	return QueueMessage{
		ProjectName: projectName,
		JobName:     name,
		JobId:       id,
		Command:     Command_Stop,
	}

}
