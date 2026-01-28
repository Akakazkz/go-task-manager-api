package service

import "errors"

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
