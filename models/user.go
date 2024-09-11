package models

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Password []byte `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Role     string `json:"role" db:"role"`
}

func (u User) IsNil() bool {
	uNil := User{}
	return (uNil.UserId == u.UserId &&
		uNil.Username == u.Username &&
		uNil.Email == u.Email &&
		uNil.Role == u.Role)
}
