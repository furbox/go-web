package utils

import (
	"regexp"
	"unicode"
)

var Regex_correo = regexp.MustCompile("^[a-z0-9._%+]+@[a-z0-9.-]+\\.[a-z]{2,4}$")

func ValidaPassword(password string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(password) >= 8 && len(password) <= 20 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
