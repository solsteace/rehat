package models

import "time"

type User struct {
	UserId     int64      `json:"user_id" db:"user_id"`
	Email      string     `json:"email" db:"email"`
	Username   string     `json:"username" db:"username"`
	Name       string     `json:"name" db:"name"`
	Password   []byte     `json:"password,omitempty" db:"password"`
	Role       string     `json:"role" db:"role"`
	IsVerified bool       `json:"verified" db:"verified"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (u User) TableName() string { return "users" }
