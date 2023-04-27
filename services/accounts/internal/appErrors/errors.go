package apperrors

import "errors"

var (
	ErrorUserNotFound      = errors.New("user doesn't exists")
	ErrUserAlreadyExists   = errors.New("user with such email already exists")
	ErrUserNotCompleteData = errors.New("All input fields are required")
)
