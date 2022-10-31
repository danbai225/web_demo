package stringutil

import (
	"crypto/sha256"
	"fmt"
)

func StringSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
