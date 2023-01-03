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
