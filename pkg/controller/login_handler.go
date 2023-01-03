package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/service"
	"github.com/hamster-shared/a-line/pkg/utils"
	"net/http"
)

func (h *HandlerServer) loginWithGithub(gin *gin.Context) {
	var loginParam parameter.LoginParam
	err := gin.BindJSON(&loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	data, err := loginService.LoginWithGithub(loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) githubRepoAuth(gin *gin.Context) {
	var authData parameter.AuthParam
	err := gin.BindJSON(&authData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	token, err := loginService.GithubRepoAuth(authData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(token, gin)
}

func (h *HandlerServer) Authorize() gin.HandlerFunc {
	return func(gin *gin.Context) {
		accessToken := gin.Request.Header.Get("Access-Token")
		if accessToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			return
		}
		token := utils.AesDecrypt(accessToken, consts.SecretKey)
		userService := application.GetBean[*service.UserService]("userService")
		user, err := userService.GetUserByToken(token)
		if err != nil {
			Failed(http.StatusUnauthorized, err.Error(), gin)
			return
		}
		if user.Token == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			return
		}
		gin.Next()
	}
}
