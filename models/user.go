package models

import "time"

// User Model
type User struct {
	ID        uint       `json:"id"`
	Username  string     `gorm:"unique_index" sql:"type:varchar(20)" json:"username"`
	Password  string     `json:"password,omitempty"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Email     string     `gorm:"unique_index" sql:"type:varchar(50)" json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
