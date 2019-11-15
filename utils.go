package main

import (
	"crypto/md5"
	"encoding/hex"
)

//md5sum - calculates md5 hash from string
func md5sum(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
