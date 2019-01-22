package util

import "regexp"

const (
	patternNumber   = "^[0-9]*$"
	patternEmail    = ""
	patternPassword = ""
)

var (
	regNumber   = regexp.MustCompile(patternNumber)
	regEmail    = regexp.MustCompile(patternEmail)
	regPassword = regexp.MustCompile(patternPassword)
)

func IsNumber(str string) bool {
	return regNumber.Match([]byte(str))
}

func IsEmail(str string) bool {
	return regEmail.Match([]byte(str))
}

func IsValidPassword(str string) bool {
	return regPassword.Match([]byte(str))
}
