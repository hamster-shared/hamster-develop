package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

// 1:github 2:metamask
func GenerateJWT(userId int, loginType int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"loginType": loginType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 24 hour
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("失败！！！！！")
		return "", err
	}
	return tokenString, nil
}
