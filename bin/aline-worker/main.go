package main

import (
	"flag"
	"time"

	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	var done = make(chan bool)
	logger.Init().ToStdoutAndFile().SetLevel(logrus.TraceLevel)
	masterAddress := parseArgs()

	for {
		_, err := engine.NewWorkerEngine(masterAddress)
		if err != nil {
			logger.Errorf("new worker engine failed, err: %v", err)
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	<-done
}

// 解析命令行参数
func parseArgs() string {
	var masterAddress string
	flag.StringVar(&masterAddress, "master", "0.0.0.0:50001", "master address")
	flag.Parse()
	return masterAddress
}
