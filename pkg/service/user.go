package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserByToken(token string) (db2.User, error)
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

func (u *UserService) GetUserCount() (int64, error) {
	var count int64
	if err := u.db.Model(&db2.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
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
