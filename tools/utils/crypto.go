package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5(s string) string {
	digest := md5.New()
	digest.Write([]byte(s))
	return hex.EncodeToString(digest.Sum(nil))
}

func Sha1(s string) string {
	digest := sha1.New()
	digest.Write([]byte(s))
	return hex.EncodeToString(digest.Sum(nil))
}

func GenPassword(pwd, salt string) string {
	return Md5(pwd + salt)
}
