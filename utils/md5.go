package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	signByte := []byte(str)
	hash := md5.New()
	hash.Write(signByte)

	return hex.EncodeToString(hash.Sum(nil))
}
