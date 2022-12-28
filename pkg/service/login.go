package service

import (
	"github.com/hamster-shared/a-line/pkg/application"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/utils"
	"gorm.io/gorm"
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
	userData.Token = token.AccessToken
	userData.AvatarUrl = *userInfo.AvatarURL
	userData.HtmlUrl = *userInfo.HTMLURL
	userData.CreateTime = time.Now()
	l.db.Create(&userData)
	return userData, nil
}
