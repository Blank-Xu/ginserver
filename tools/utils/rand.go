package utils

import (
	"bytes"
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	defaultLen     = 32
	defaultSafeLen = 16

	defaultSeedIndex = 3
)

var (
	seed1 = []byte("0123456789")
	seed2 = []byte("abcdefghijklmnopqrstuvwxyz")
	seed3 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	seed4 = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	seed5 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	seed6 = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	seed7 = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	seed8 = append(seed4, []byte("~!@#$%^&*()_+")...)

	seeds = [8][]byte{seed1, seed2, seed3, seed4, seed5, seed6, seed7, seed8}
)

func randSafeStr(length int) (string, error) {
	b := make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GenRandSafeStr(length int) (string, error) {
	return randSafeStr(length)
}

func GenSalt() (salt string) {
	salt, _ = randSafeStr(defaultSafeLen)
	return
}

func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randStr(length, seedIndex int) string {
	r := NewRand()
	seed := seeds[seedIndex]
	seedLen := len(seed)
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		buf.WriteByte(seed[r.Intn(seedLen)])
	}
	return buf.String()
}

func GenRandStr(length int) string {
	return randStr(length, defaultSeedIndex)
}
