package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
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

func (h *HandlerServer) githubWebHook(gin *gin.Context) {
	var githubWebHook parameter.GithubWebHook
	err := gin.BindJSON(&githubWebHook)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	event := gin.GetHeader("X-GitHub-Event")
	if event == "installation" && githubWebHook.Action == "deleted" {
		userService := application.GetBean[*service.UserService]("userService")
		user, err := userService.GetUserById(githubWebHook.Sender.Id)
		if err == nil {
			user.Token = ""
			userService.UpdateUser(user)
		}
	}
}

func (h *HandlerServer) githubInstall(gin *gin.Context) {
	var installData parameter.InstallParam
	err := gin.BindJSON(&installData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	token, err := loginService.GithubInstall(installData.Code)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(token, gin)
}

func (h *HandlerServer) RequestLog() gin.HandlerFunc {
	return func(gin *gin.Context) {
		accessToken := gin.Request.Header.Get("Access-Token")
		url := gin.Request.RequestURI
		log.Printf("url: %s, method: %s, token: %s ", url, gin.Request.Method, accessToken)
		requestLog := &db2.RequestLog{
			Url:        url,
			Token:      accessToken,
			Method:     gin.Request.Method,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		db := application.GetBean[*gorm.DB]("db")
		_ = db.Save(requestLog).Error
		gin.Next()
	}
}

func (h *HandlerServer) Authorize() gin.HandlerFunc {
	return func(gin *gin.Context) {
		accessToken := gin.Request.Header.Get("Access-Token")
		if accessToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			gin.Abort()
			return
		}
		if !strings.HasPrefix(accessToken, "0x") {
			token := utils.AesDecrypt(accessToken, consts.SecretKey)
			userService := application.GetBean[*service.UserService]("userService")
			user, err := userService.GetUserByToken(token)
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				gin.Abort()
				return
			}
			if user.Token == "" {
				Failed(http.StatusUnauthorized, "access not authorized", gin)
				gin.Abort()
				return
			}
			user.Token = accessToken
			gin.Set("token", token)
			gin.Set("user", user)
		}
		gin.Next()
	}
}

func (h *HandlerServer) getUseInfo(gin *gin.Context) {
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	user, _ := userAny.(db2.User)
	Success(user, gin)
}

func (h *HandlerServer) getUserCount(gin *gin.Context) {
	userService := application.GetBean[*service.UserService]("userService")
	data, err := userService.GetUserCount()
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) saveUserWallet(gin *gin.Context) {
	var userId uint
	userAny, exit := gin.Get("user")

	if exit {
		user, _ := userAny.(db2.User)
		userId = user.Id
	}
	userService := application.GetBean[*service.UserService]("userService")
	wallet := &db2.UserWallet{}
	err := gin.BindJSON(wallet)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	userService.SaveUserWallet(userId, wallet.Address)
	Success("", gin)

}

func (h *HandlerServer) updateFirstState(gin *gin.Context) {
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	user, _ := userAny.(db2.User)
	if user.FirstState == 0 {
		userService := application.GetBean[*service.UserService]("userService")
		userData, err := userService.GetUserById(int64(user.Id))
		if err != nil {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			return
		}
		userData.FirstState = 1
		err = userService.UpdateUser(userData)
		if err != nil {
			Fail(err.Error(), gin)
			return
		}
		gin.Set("user", user)
	}
	Success("", gin)
}
