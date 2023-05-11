package services

import (
	"errors"
	"regexp"
)

func validateCredentials(email, password string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("incorrect email format")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 8 || !regexp.MustCompile(`[A-Z]+`).MatchString(password) || !regexp.MustCompile(`\d+`).MatchString(password) {
		return errors.New("Password must be at least 8 characters long and contain at least one uppercase letter and one digit")
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 8 || !regexp.MustCompile(`[A-Z]+`).MatchString(password) || !regexp.MustCompile(`\d+`).MatchString(password) {
		return errors.New("Password must be at least 8 characters long and contain at least one uppercase letter and one digit")
	}
	return nil
}
