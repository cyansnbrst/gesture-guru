package db

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user with this email already exists")
)
