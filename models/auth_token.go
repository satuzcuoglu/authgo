package models

import (
	"time"
)

// AuthToken keeps user and token info
type AuthToken struct {
	ID        uint      `json:"-"`
	Token     string    `json:"token"`
	User      User      `json:"user" gorm:"association_autoupdate:false"`
	UserID    uint      `json:"-"`
	CreatedAt time.Time `json:"-"`
	ExpiresAt time.Time `json:"-"`
}
