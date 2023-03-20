package service

import (
	"errors"
	"fmt"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type ILoginService interface {
}

type LoginService struct {
	db            *gorm.DB
	githubService *GithubService
}

func NewLoginService() *LoginService {
	return &LoginService{
		db:            application.GetBean[*gorm.DB]("db"),
		githubService: application.GetBean[*GithubService]("githubService"),
	}
}

func (l *LoginService) LoginWithGithub(data parameter.LoginParam) (vo.UserVo, error) {
	//data.ClientSecret = consts.ClientSecrets
	data.ClientSecret = os.Getenv("CLIENT_SECRETS")
	var userData db2.User
	var userVo vo.UserVo
	var token parameter.Token
	url := "https://github.com/login/oauth/access_token"
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"client_id":     data.ClientId,
		"client_secret": data.ClientSecret,
		"code":          data.Code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if res.StatusCode() != 200 {
		return userVo, err
	}
	if err != nil {
		return userVo, err
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		return userVo, err
	}
	err = l.db.Model(db2.User{}).Where("id = ?", userInfo.ID).First(&userData).Error
	if err != nil {
		userData.Id = uint(*userInfo.ID)
		userData.Username = *userInfo.Login
		userData.AvatarUrl = *userInfo.AvatarURL
		userData.HtmlUrl = *userInfo.HTMLURL
		userData.CreateTime = time.Now()
		l.db.Save(&userData)
	}
	copier.Copy(&userVo, &userData)
	if userData.Token != "" {
		accessToken := utils.AesEncrypt(userData.Token, consts.SecretKey)
		userVo.Token = accessToken
	}
	return userVo, nil
}

func (l *LoginService) GithubInstall(code string) (string, error) {
	var userData db2.User
	var token parameter.Token
	url := "https://github.com/login/oauth/access_token"
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		//"client_id":     consts.AppsClientId,
		"client_id": os.Getenv("APPS_CLIENT_ID"),
		//"client_secret": consts.AppsClientSecrets,
		"client_secret": os.Getenv("APPS_CLIENT_SECRETS"),
		"code":          code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if res.StatusCode() != 200 {
		return "", err
	}
	if err != nil {
		return "", err
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		log.Println("github install failed:user not found", err.Error())
		return "", err
	}
	email, err := l.githubService.GetUserEmail(token.AccessToken)
	if err != nil {
		log.Println("github install failed:get email failed", err.Error())
		return "", err
	}
	err = l.db.Model(db2.User{}).Where("id = ?", userInfo.ID).First(&userData).Error
	if err != nil {
		return "", err
	}
	userData.UserEmail = email
	userData.Token = token.AccessToken
	l.db.Save(&userData)
	accessToken := utils.AesEncrypt(token.AccessToken, consts.SecretKey)
	return accessToken, nil
}

func (l *LoginService) GithubRepoAuth(authData parameter.AuthParam) (string, error) {
	//authData.ClientSecret = consts.ClientSecrets
	authData.ClientSecret = os.Getenv("CLIENT_SECRETS")
	var userData db2.User
	var token parameter.Token
	res := l.db.Model(db2.User{}).Where("id = ?", authData.UserId).First(&userData)
	if res.Error != nil {
		log.Println("login user not fond ", res.Error)
		return "", res.Error
	}

	url := "https://github.com/login/oauth/access_token"
	response, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"client_id":     authData.ClientId,
		"client_secret": authData.ClientSecret,
		"code":          authData.Code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if response.StatusCode() != 200 {
		log.Println(string(response.Body()))
		return "", errors.New(fmt.Sprintf("auth failed:%s", string(response.Body())))
	}
	if err != nil {
		log.Println("repo auth failed ", err)
		return "", err
	}
	userData.Token = token.AccessToken
	l.db.Save(&userData)
	accessToken := utils.AesEncrypt(token.AccessToken, consts.SecretKey)
	return accessToken, nil
}
