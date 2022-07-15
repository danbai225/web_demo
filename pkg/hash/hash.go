package hash

import (
	"crypto/md5"
	"encoding/hex"
)

var _ Hash = (*hash)(nil)

type Hash interface {
	i()

	// HashidsEncode 加密
	HashidsEncode(params []int) (string, error)

	// HashidsDecode 解密
	HashidsDecode(hash string) ([]int, error)
}

type hash struct {
	secret string
	length int
}

func New(secret string, length int) Hash {
	return &hash{
		secret: secret,
		length: length,
	}
}

func (h *hash) i() {}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
func MD5Byte(s []byte) string {
	sum := md5.Sum(s)
	return hex.EncodeToString(sum[:])
}
