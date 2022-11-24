package token

import (
	"crypto/sha1"
	"fmt"
)

func Hash(s string) string {
	sha := sha1.New()
	sha.Write([]byte(s))
	encrypted := sha.Sum(nil)
	h := fmt.Sprintf("%x", encrypted)
	return h
}
