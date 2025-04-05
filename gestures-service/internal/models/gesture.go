package models

import (
	"time"
)

// Gesture struct
type Gesture struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VideoURL    string    `json:"video_url"`
	CategoryID  Category  `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}
