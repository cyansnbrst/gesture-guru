package models

import "time"

// User struct
type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	IsAdmin      bool
	CreatedAt    time.Time
}
