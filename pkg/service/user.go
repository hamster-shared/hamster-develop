package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserByToken(token string) (db2.User, error)
	GetGithubInstallId(userId uint) (int64, error)
}

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (u *UserService) GetUserByToken(token string) (db2.User, error) {
	var user db2.User
	res := u.db.Model(db2.User{}).Where("token = ?", token).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserService) UpdateUser(user db2.User) error {
	return u.db.Save(&user).Error
}

func (u *UserService) GetUserById(id int64) (db2.User, error) {
	var user db2.User
	res := u.db.Model(db2.User{}).Where("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserService) GetGithubInstallId(userId uint) (int64, error) {
	var gitAppInstall db2.GitAppInstall
	err := u.db.Model(db2.GitAppInstall{}).Where("user_id = ?", userId).First(&gitAppInstall).Error
	if err != nil {
		return 0, err
	}
	return gitAppInstall.InstallId, nil
}

func (u *UserService) GetUserCount() (int64, error) {
	var count int64
	if err := u.db.Model(&db2.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserService) GetGithubUser(id int64) (db2.User, error) {
	var user db2.User
	err := u.db.Model(db2.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *UserService) SaveUserWallet(userId uint, address string) int64 {

	var count int64
	err := u.db.Model(&db2.UserWallet{}).Where("address=?", address).Count(&count).Error
	if err == nil && count == 0 {
		wallet := &db2.UserWallet{
			UserId:  userId,
			Address: address,
		}
		err := u.db.Save(wallet).Error
		if err != nil {
			return 0
		}
	}
	return count
}

func (u *UserService) GetUserWalletById(id int) (db2.UserWallet, error) {
	var userWallet db2.UserWallet
	err := u.db.Model(db2.UserWallet{}).Where("id = ?", id).First(&userWallet).Error
	return userWallet, err
}

func (u *UserService) UpdateUserWallet(userWallet db2.UserWallet) error {
	return u.db.Save(&userWallet).Error
}

func (u *UserService) SaveUserToken(id uint, token string) error {
	return u.db.Model(db2.User{}).Where("id = ?", id).Update("token", token).Error
}
