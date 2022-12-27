package service

import (
	"github.com/hamster-shared/a-line/pkg/application"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/utils"
	"gorm.io/gorm"
	"log"
)

type ILoginService interface {
}

type LoginService struct {
	db *gorm.DB
}

func NewLoginService() *LoginService {
	return &LoginService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (l *LoginService) LoginWithGithub(data parameter.LoginParam) (db2.User, error) {
	var userData db2.User
	req := utils.NewHttp().NewRequest()
	url := "https://github.com/login/oauth/access_token"
	req.SetPathParam("client_id", data.ClientId)
	req.SetPathParam("client_secret", data.ClientSecret)
	req.SetPathParam("code", data.Code)
	resp, err := req.Post(url)
	if err != nil {
		return userData, err
	}
	log.Println(resp)
	return userData, nil
}
