package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(str string) string {
	hash := md5.Sum([]byte(str))
	strHash := hex.EncodeToString(hash[:])
	return strHash
}
