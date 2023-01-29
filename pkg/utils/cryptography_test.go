package utils

import (
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"log"
	"testing"
)

func Test_Cryptography(t *testing.T) {
	data := "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"
	str := AesEncrypt(data, consts.SecretKey)
	log.Println(str)
	token := AesDecrypt(str, consts.SecretKey)
	log.Println(token)
}
