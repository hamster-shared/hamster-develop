package utils

import (
	"github.com/hamster-shared/a-line/pkg/consts"
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
