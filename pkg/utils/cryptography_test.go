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
