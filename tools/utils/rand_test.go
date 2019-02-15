package utils

import (
	"fmt"
	"testing"
)

func TestGenSalt(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(GenSalt())
	}
}

func BenchmarkGenSalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenSalt()
	}
}

func TestGenRandStr(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(GenRandStr(32))
	}
}

func BenchmarkGenRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandStr(32)
	}
}
