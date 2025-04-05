package db

import "errors"

var (
	ErrGestureNotFound = errors.New("gesture not found")
	ErrInvalidCategory = errors.New("invalid category number")
)
