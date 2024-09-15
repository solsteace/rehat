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

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE username=?",
		user.TableName())
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
		&user.IsVerified,
		&user.DeletedAt)
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

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id=?",
		user.TableName())
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
		&user.IsVerified,
		&user.DeletedAt)
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

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE email=?",
		user.TableName())
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
		&user.IsVerified,
		&user.DeletedAt)
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
	query := fmt.Sprintf(
		`INSERT 
			INTO %s(username, name, password, email, role) 
			VALUES (?, ?, ?, ?, ?)`,
		user.TableName())
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
		user.Role)
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
	query := fmt.Sprintf(
		`UPDATE %s 
			SET email = ?,
				username = ?,
				name = ?,
				password = ?,
				role = ?
			WHERE id = ?`,
		user.TableName())
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
