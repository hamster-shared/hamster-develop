package main

import (
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/hamster-develop/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.Init().ToStdoutAndFile().SetLevel(logrus.DebugLevel)
	cmd.Execute()
}
