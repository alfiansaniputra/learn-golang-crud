package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int          `json:"id"`
	Username      string       `json:"username"`
	Photo         string       `json:"photo"`
	Password      string       `json:"password"`
	Email         string       `json:"email"`
	FCMToken      string       `json:"fcm_token"`
	RememberToken string       `json:"remember_token"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

func (u *User) SetDefaults() {
	if u.Username == "" {
		u.Username = "-"
	}
	if u.Photo == "" {
		u.Photo = "-"
	}
	if u.Password == "" {
		u.Password = "-"
	}
	if u.Email == "" {
		u.Email = "-"
	}
	if u.FCMToken == "" {
		u.FCMToken = "-"
	}
	if u.RememberToken == "" {
		u.RememberToken = "-"
	}
}
