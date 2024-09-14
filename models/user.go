package models

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Password []byte `json:"password,omitempty" db:"password"`
	Email    string `json:"email" db:"email"`
	Role     string `json:"role" db:"role"`
	IsActive bool   `json:"is_active" db:"active"`
}
