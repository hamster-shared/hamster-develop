package main

import (
	"github.com/hamster-shared/node-api/cmd"
	"github.com/hamster-shared/node-api/pkg/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	logger.Init().ToStdoutAndFile().SetLevel(logrus.TraceLevel)
	cmd.Execute()
}
