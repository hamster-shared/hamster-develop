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

func (l *LoginService) LoginWithGithubV2(data parameter.LoginParam) (string, error) {
	var userData db2.User
	var token parameter.Token
	url := "https://github.com/login/oauth/access_token"
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"client_id":     os.Getenv("APPS_CLIENT_ID"),
		"client_secret": os.Getenv("APPS_CLIENT_SECRETS"),
		//"client_id":     "Iv1.c41a1e51c5ebcf42",
		//"client_secret": "419540e126f38890c3974d9b082de63324fa0be8",
		"code": data.Code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if res.StatusCode() != 200 {
		return "", err
	}
	if err != nil {
		return "", err
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		return "", err
	}
	email, err := l.githubService.GetUserEmail(token.AccessToken)
	if err != nil {
		log.Println("github install failed:get email failed", err.Error())
		return "", err
	}
	err = l.db.Model(db2.User{}).Where("id = ?", userInfo.ID).First(&userData).Error
	userData.UserEmail = email
	if err != nil {
		userData.Id = uint(*userInfo.ID)
		userData.Username = *userInfo.Login
		userData.AvatarUrl = *userInfo.AvatarURL
		userData.HtmlUrl = *userInfo.HTMLURL
		userData.CreateTime = time.Now()
		userData.LoginType = consts.GitHub
		l.db.Create(&userData)
	} else {
		l.db.Save(&userData)
	}
	jwtToken, err := utils.GenerateJWT(int(userData.Id), consts.GitHub)
	return jwtToken, err
}

// true:need install false: not need install
func (l *LoginService) GithubInstallAuth(data parameter.LoginParam, userWallet db2.UserWallet) (bool, error) {
	data.ClientSecret = os.Getenv("APPS_CLIENT_SECRETS")
	var userData db2.User
	var token parameter.Token
	result := false
	url := "https://github.com/login/oauth/access_token"
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"client_id":     data.ClientId,
		"client_secret": data.ClientSecret,
		"code":          data.Code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if err != nil {
		return result, err
	}
	if res.IsError() {
		return result, errors.New(res.String())
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		return result, err
	}
	err = l.db.Model(db2.User{}).Where("id = ?", userInfo.ID).First(&userData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = true
		userData.Id = uint(userInfo.GetID())
		userData.Username = userInfo.GetLogin()
		userData.LoginType = consts.Metamask
		userData.AvatarUrl = userInfo.GetAvatarURL()
		userData.FirstState = 0
		userData.HtmlUrl = userInfo.GetHTMLURL()
		userData.CreateTime = time.Now()
		l.db.Model(db2.User{}).Create(&userData)
	}
	userWallet.UserId = userData.Id
	l.db.Save(&userWallet)
	result = true
	return result, nil
}

func (l *LoginService) MetamaskLogin(data parameter.MetaMaskLoginParam) (string, error) {
	var userWallet db2.UserWallet
	err := l.db.Model(db2.UserWallet{}).Where("address = ?", data.Address).First(&userWallet).Error
	if err != nil {
		userWallet.CreateTime = time.Now()
		userWallet.Address = data.Address
		l.db.Save(&userWallet)
	}
	jwtToken, err := utils.GenerateJWT(int(userWallet.Id), consts.Metamask)
	return jwtToken, err
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

func (l *LoginService) GithubInstallV2(code string, loginType int, userWallet db2.UserWallet) error {
	var userData db2.User
	var token parameter.Token
	url := "https://github.com/login/oauth/access_token"
	_, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		//"client_id":     consts.AppsClientId,
		"client_id": os.Getenv("APPS_CLIENT_ID"),
		//"client_secret": consts.AppsClientSecrets,
		"client_secret": os.Getenv("APPS_CLIENT_SECRETS"),
		"code":          code,
	}).SetResult(&token).SetHeader("Accept", "application/json").Post(url)
	if err != nil {
		return err
	}
	userInfo, err := l.githubService.GetUserInfo(token.AccessToken)
	if err != nil {
		log.Println("github install v2 failed:user not found", err.Error())
		return err
	}
	email, err := l.githubService.GetUserEmail(token.AccessToken)
	if err != nil {
		log.Println("v2 github install failed:get email failed", err.Error())
		return err
	}
	err = l.db.Model(db2.User{}).Where("id = ?", userInfo.ID).First(&userData).Error
	if err != nil {
		if loginType == consts.GitHub {
			log.Println("v2 user not found!")
			return err
		}
		userData.Id = uint(*userInfo.ID)
		userData.Username = *userInfo.Login
		userData.AvatarUrl = *userInfo.AvatarURL
		userData.HtmlUrl = *userInfo.HTMLURL
		userData.CreateTime = time.Now()
		userData.LoginType = consts.GitHub
	}
	userData.Token = token.AccessToken
	userData.UserEmail = email
	l.db.Save(&userData)
	if loginType == consts.Metamask {
		userWallet.UserId = userData.Id
		l.db.Save(&userWallet)
	}
	return nil
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
