package models

import (
	"encoding/json"
	"time"
)

// User Model
type User struct {
	ID        uint       `json:"id" binding:"required"`
	Username  string     `gorm:"unique_index" sql:"type:varchar(20)" json:"username" binding:"required"`
	Password  string     `json:"password,omitempty" binding:"required"`
	Firstname string     `json:"firstname" binding:"required"`
	Lastname  string     `json:"lastname" binding:"required"`
	Email     string     `gorm:"unique_index" sql:"type:varchar(50)" json:"email" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// MarshalJSON used for Marshaling User
func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}
