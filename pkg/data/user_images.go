package data

import "time"

// UserImages is type for user profile images struct
type UserImages struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
