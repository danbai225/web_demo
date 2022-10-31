package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"web_demo/internal/myconst"
	"web_demo/pkg/stringutil"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var c cipher.Block

func init() {
	var err error
	c, err = aes.NewCipher([]byte("web_demoweb_demo"))
	if err != nil {
		panic(err)
	}
}
func StringHash(str string) string {
	return stringutil.StringSha256(fmt.Sprint(str, myconst.PasswordSalt))
}

func Encryption(plaintext string) (string, error) {
	ciphertext := make([]byte, len(plaintext))
	cipher.NewCFBEncrypter(c, commonIV).XORKeyStream(ciphertext, []byte(plaintext))
	return string(ciphertext), nil

}
func Decrypt(plaintext string) string {
	plaintextCopy := make([]byte, len(plaintext))
	cipher.NewCFBDecrypter(c, commonIV).XORKeyStream(plaintextCopy, []byte(plaintext))
	return string(plaintextCopy)
}
