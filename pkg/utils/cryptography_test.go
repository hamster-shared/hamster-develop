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
