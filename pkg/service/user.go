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
