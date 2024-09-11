package services

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type User struct {
	Db *sql.DB
}

func (u User) Create(newUser models.User) (int64, error) {
	query := `INSERT INTO users(username, name, password, email) 
				VALUES (?, ?, ?, ?)`
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		newUser.Username,
		newUser.Name,
		newUser.Password,
		newUser.Email)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (u User) GetByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username=?"
	stmt, err := u.Db.Prepare(query)
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrSQLNoRows{
				message: fmt.Sprintf("Couldn't find user with username %s", username)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}

func (u User) GetById(id string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=?"
	stmt, err := u.Db.Prepare(query)
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrSQLNoRows{
				message: fmt.Sprintf("Couldn't find user with id %s", id)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}

func (u User) GetByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email=?"
	stmt, err := u.Db.Prepare(query)
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &ErrSQLNoRows{
				message: fmt.Sprintf("Couldn't find user with email %s", email)}
		}
		return user, &ErrSQL{message: err.Error()}
	}

	return user, nil
}
