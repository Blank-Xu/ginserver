package utils

import (
	crand "crypto/rand"
	"encoding/hex"
)

const (
	lenSalt = 32
)

func GenSafeRandomStr(length int) (string, error) {
	b := make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

const (
	seed1 = "0123456789"
	seed2 = "abcdefghijklmnopqrstuvwxyz"
	seed3 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed4 = "0123456789abcdefghijklmnopqrstuvwxyz"
	seed5 = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	seed6 = "abcdefghjklmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	seed7 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//func GenRandomStr(l int) (string, error) {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//}
