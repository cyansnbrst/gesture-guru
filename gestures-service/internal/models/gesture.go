package models

import (
	"time"
)

// Gesture struct
type Gesture struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VideoURL    string    `json:"video_url"`
	CategoryID  int64     `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}
