package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func NameValidator(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	if len(name) < 3 || len(name) > 200 {
		return false
	}

	for _, char := range name {
		if !(char == ' ' || char == '-' || unicode.IsLetter(char) || unicode.IsDigit(char)) {
			return false
		}
	}

	return true
}

func LocationValidator(fl validator.FieldLevel) bool {
	location := fl.Field().String()

	if len(location) < 3 || len(location) > 50 {
		return false
	}

	for _, char := range location {
		if !(char == ' ' || char == '-' || unicode.IsLetter(char)) {
			return false
		}
	}

	return true
}

func CountryValidator(fl validator.FieldLevel) bool {
	country := fl.Field().String()

	if len(country) < 3 || len(country) > 50 {
		return false
	}

	for _, char := range country {
		if !(char == ' ' || char == '-' || unicode.IsLetter(char)) {
			return false
		}
	}

	return true
}

func DescriptionValidator(fl validator.FieldLevel) bool {
	description := fl.Field().String()

	if len(description) < 10 {
		return false
	}

	return true
}

func UsernameValidator(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	if len(username) < 3 || len(username) > 40 {
		return false
	}

	for _, char := range username {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_' || char == '-' || char == '.') {
			return false
		}
	}

	return true
}

func PasswordValidator(fl validator.FieldLevel) bool {
	if password, ok := fl.Field().Interface().(string); ok {
		return isComplexPassword(password)
	}
	return false
}

func isComplexPassword(password string) bool {
	var (
		hasMinLen  = len(password) >= 8
		hasMaxLen  = len(password) <= 40
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		print(char)
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune("@$!%*?&", char):

			print("Special")
			hasSpecial = true
		}
	}
	print(hasMinLen, hasMaxLen, hasUpper, hasLower, hasNumber, hasSpecial)

	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}
