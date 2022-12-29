package service

import (
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/consts"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/utils"
	"gorm.io/gorm"
	"log"
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

func (l *LoginService) LoginWithGithub(data parameter.LoginParam) (db2.User, error) {
	var userData db2.User
	var token parameter.Token
	req := utils.NewHttp().NewRequest()
	url := "https://github.com/login/oauth/access_token"
	req.SetPathParam("client_id", data.ClientId)
	req.SetPathParam("client_secret", data.ClientSecret)
	req.SetPathParam("code", data.Code)
	_, err := req.SetResult(&token).Post(url)
	if err != nil {
		return userData, err
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		return userData, err
	}
	userData.Id = uint(*userInfo.ID)
	userData.Username = *userInfo.Name
	userData.AvatarUrl = *userInfo.AvatarURL
	userData.HtmlUrl = *userInfo.HTMLURL
	userData.CreateTime = time.Now()
	l.db.Save(&userData)
	return userData, nil
}

func (l *LoginService) GithubRepoAuth(authData parameter.AuthParam) (string, error) {
	var userData db2.User
	var token parameter.Token
	res := l.db.Model(db2.User{}).Where("id = ?", authData.UserId).First(&userData)
	if res.Error != nil {
		log.Println("login user not fond ", res.Error)
		return "", res.Error
	}
	req := utils.NewHttp().NewRequest()
	url := "https://github.com/login/oauth/access_token"
	req.SetPathParam("client_id", authData.ClientId)
	req.SetPathParam("client_secret", authData.ClientSecret)
	req.SetPathParam("code", authData.Code)
	_, err := req.SetResult(&token).Post(url)
	if err != nil {
		log.Println("repo auth failed ", err)
		return "", err
	}
	userData.Token = token.AccessToken
	l.db.Save(&userData)
	accessToken := utils.AesEncrypt(token.AccessToken, consts.SecretKey)
	return accessToken, nil
}
