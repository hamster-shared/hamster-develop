package controller

import (
	"testing"

	"github.com/hamster-shared/aline-engine/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_getGithubRawUrl(t *testing.T) {
	logger.Init().ToStdoutAndFile().SetLevel(logrus.TraceLevel)
	rawUrl := getGithubRawUrl("https://github.com/hamster-template/aptos-raffle", "main", "Move.toml")
	assert.Equal(t, "https://raw.githubusercontent.com/hamster-template/aptos-raffle/main/Move.toml", rawUrl)
}
