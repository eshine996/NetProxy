package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 string to 32-bit MD5 strings
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
