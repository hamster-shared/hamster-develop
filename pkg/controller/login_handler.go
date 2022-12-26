package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/parameter"
)

func (h *HandlerServer) loginWithGithub(gin *gin.Context) {
	var loginParam parameter.LoginParam
	err := gin.BindJSON(&loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
}
