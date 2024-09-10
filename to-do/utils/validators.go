package utils

import (
	"errors"
	"strings"
)

func PasswordValidator(password string) error {
	upperCase, lowercase, number := 0, 0, 0
	if len(password) > 8 {
		for _, i := range password {
			// Uppercase check
			if i >= 64 && i <= 90 {
				upperCase += 1
			}
			// Lowercase check
			if i >= 97 && i <= 122 {
				lowercase += 1
			}
			// Integer check
			if i >= '0' && i <= '9' {
				number += 1
			}
		}
		if upperCase == 0 || lowercase == 0 || number == 0 {
			return errors.New("invalid password")
		}
	} else {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func EmailValidator(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return errors.New("invalid email")
}
