package utils

import (
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"log"
	"testing"
)

func Test_Cryptography(t *testing.T) {
	data := "ghu_pf5Q8Os00cQD4dUfDcQ755mayhlWxA4G26fu"
	str := AesEncrypt(data, consts.SecretKey)
	log.Println(str)
	token := AesDecrypt(str, consts.SecretKey)
	log.Println(token)
}
