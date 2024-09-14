package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Password []byte `json:"password,omitempty" db:"password"`
	Email    string `json:"email" db:"email"`
	Role     string `json:"role" db:"role"`
	IsActive bool   `json:"is_active" db:"active"`
}

func (u User) IsNil() bool {
	uNil := User{}
	return (uNil.UserId == u.UserId &&
		uNil.Username == u.Username &&
		uNil.Email == u.Email &&
		uNil.Role == u.Role)
}

func (u User) GetByUsername(db *sql.DB, username string) (User, error) {
	var user User

	query := "SELECT * FROM users WHERE username=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}

	err = stmt.QueryRow(username).Scan(
		&user.UserId,
		&user.Email,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find user with username %s", username)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}

func (u User) GetById(db *sql.DB, id int64) (User, error) {
	var user User

	query := "SELECT * FROM users WHERE user_id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}

	err = stmt.QueryRow(id).Scan(
		&user.UserId,
		&user.Email,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find user with id %d", id)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}

func (u User) GetByEmail(db *sql.DB, email string) (User, error) {
	var user User

	query := "SELECT * FROM users WHERE email=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}

	err = stmt.QueryRow(email).Scan(
		&user.UserId,
		&user.Email,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find user with email %s", email)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}

func (u User) Save(db *sql.DB) (int64, error) {
	query := `INSERT INTO users(username, name, password, email, role, active) 
				VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		u.Username,
		u.Name,
		u.Password,
		u.Email,
		u.Role,
		u.IsActive)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (u User) EditById(db *sql.DB, id int64) (int64, error) {
	query := `UPDATE users VALUES (?, ?, ?, ?, ?, ?) WHERE id=?`
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		u.Username,
		u.Name,
		u.Password,
		u.Email,
		u.Role,
		u.IsActive,
		u.UserId)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}
