package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

func (h *HandlerServer) githubLogin(gin *gin.Context) {
	var loginParam parameter.LoginParam
	err := gin.BindJSON(&loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	data, err := loginService.LoginWithGithubV2(loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) githubInstallAuth(gin *gin.Context) {
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	user, ok := userAny.(db2.UserWallet)
	if !ok {
		Failed(http.StatusBadRequest, "the user login method is incorrect", gin)
		return
	}
	var loginParam parameter.LoginParam
	err := gin.BindJSON(&loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	data, err := loginService.GithubInstallAuth(loginParam, user)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) metamaskLogin(gin *gin.Context) {
	var loginParam parameter.MetaMaskLoginParam
	err := gin.BindJSON(&loginParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	data, err := loginService.MetamaskLogin(loginParam)
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

func (h *HandlerServer) githubInstallV2(gin *gin.Context) {
	loginType, exit := gin.Get("loginType")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	loginMethod, _ := loginType.(int)
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	var userWallet db2.UserWallet
	if loginType == consts.Metamask {
		userWallet, _ = userAny.(db2.UserWallet)
	}
	var installData parameter.InstallParam
	err := gin.BindJSON(&installData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	loginService := application.GetBean[*service.LoginService]("loginService")
	err = loginService.GithubInstallV2(installData.Code, loginMethod, userWallet)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
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

func (h *HandlerServer) JwtAuthorize() gin.HandlerFunc {
	return func(gin *gin.Context) {
		jwtToken := gin.GetHeader("Authorization")
		log.Println(jwtToken)
		if jwtToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			gin.Abort()
			return
		}
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		userId, ok := claims["userId"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		loginType, ok := claims["loginType"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		log.Println(loginType)
		gin.Set("loginType", int(loginType))
		githubToken := ""
		userService := application.GetBean[*service.UserService]("userService")
		if loginType == consts.GitHub {
			user, err := userService.GetUserById(int64(userId))
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				gin.Abort()
				return
			}
			githubToken = user.Token
			gin.Set("user", user)
		}
		if loginType == consts.Metamask {
			userWallet, err := userService.GetUserWalletById(int(userId))
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				gin.Abort()
				return
			}
			if userWallet.UserId != 0 {
				user, err := userService.GetUserById(int64(userId))
				if err != nil {
					logger.Errorf("wallet user id is error: %s", err)
					Failed(http.StatusUnauthorized, err.Error(), gin)
					gin.Abort()
					return
				}
				githubToken = user.Token
				gin.Set("user", user)
			} else {
				gin.Set("user", userWallet)
			}
		}
		gin.Set("token", githubToken)
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

func (h *HandlerServer) getUseInfoV2(gin *gin.Context) {
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	loginType, exit := gin.Get("loginType")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	var user interface{}
	if loginType == consts.GitHub {
		user, _ = userAny.(db2.User)
	}
	if loginType == consts.Metamask {
		user, _ = userAny.(db2.UserWallet)
	}
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

func (h *HandlerServer) updateFirstStateV2(gin *gin.Context) {
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	loginType, exit := gin.Get("loginType")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	userService := application.GetBean[*service.UserService]("userService")
	if loginType == consts.GitHub {
		user, _ := userAny.(db2.User)
		if user.FirstState == 0 {
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
	}
	if loginType == consts.Metamask {
		user, _ := userAny.(db2.UserWallet)
		if user.FirstState == 0 {
			userData, err := userService.GetUserWalletById(int(user.Id))
			if err != nil {
				Failed(http.StatusUnauthorized, "access not authorized", gin)
				return
			}
			userData.FirstState = 1
			err = userService.UpdateUserWallet(userData)
			if err != nil {
				Fail(err.Error(), gin)
				return
			}
			gin.Set("user", user)
		}
	}
	Success("", gin)
}

func (h *HandlerServer) githubInstallCheck(gin *gin.Context) {
	loginType, exit := gin.Get("loginType")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	userAny, exit := gin.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", gin)
		return
	}
	result := true
	userService := application.GetBean[*service.UserService]("userService")
	if loginType == consts.GitHub {
		user, _ := userAny.(db2.User)
		userInfo, err := userService.GetUserById(int64(user.Id))
		if err != nil {
			Failed(http.StatusUnauthorized, err.Error(), gin)
			return
		}
		if userInfo.Token == "" {
			result = false
		}
	}
	if loginType == consts.Metamask {
		user, _ := userAny.(db2.UserWallet)
		userInfo, err := userService.GetUserWalletById(int(user.Id))
		if err != nil {
			Failed(http.StatusUnauthorized, err.Error(), gin)
			return
		}
		if userInfo.UserId == 0 {
			result = false
		}
	}
	Success(result, gin)
}
