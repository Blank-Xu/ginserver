package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5Byte(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(s string) string {
	return Md5Byte([]byte(s))
}

func Sha1Byte(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(s string) string {
	return Sha1Byte([]byte(s))
}
