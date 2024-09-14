package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type User struct {
	Db *sql.DB
}

func (u User) GetByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username=?"
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

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

func (u User) GetById(id int64) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE user_id=?"
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

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

func (u User) GetByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email=?"
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return user, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

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

func (u User) Save(user models.User) (int64, error) {
	query := `INSERT INTO users(username, name, password, email, role, active) 
				VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		user.Username,
		user.Name,
		user.Password,
		user.Email,
		user.Role,
		user.IsActive)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (u User) EditById(id int64, user models.User) (int64, error) {
	query := `UPDATE users VALUES (?, ?, ?, ?, ?, ?) WHERE id=?`
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		user.Username,
		user.Name,
		user.Password,
		user.Email,
		user.Role,
		user.IsActive,
		user.UserId)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}
