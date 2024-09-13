package models

import (
	"database/sql"
	"fmt"
)

type MotelAdmin struct {
	AdminID int64 `json:"admin_id" db:"admin_id"`
	UserID  int64 `json:"user_id" db:"user_id"`
	MotelID int64 `json:"motel_id" db:"motel_id"`
}

func (m MotelAdmin) Save(db *sql.DB) (int64, error) {
	query := `INSERT INTO motel_admins(user_id, motel_id) VALUES (?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		m.UserID,
		m.MotelID)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (m MotelAdmin) GetById(db *sql.DB, id string) (MotelAdmin, error) {
	var motel MotelAdmin

	query := "SELECT * FROM motel_admins WHERE admin_id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return motel, &ErrSQL{message: err.Error()}
	}

	err = stmt.QueryRow(id).Scan(
		&motel.AdminID,
		&motel.UserID,
		&motel.MotelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return motel, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find motel admin with id %s", id)}
		}
		return motel, &ErrSQL{message: err.Error()}
	}
	return motel, nil
}

func (m MotelAdmin) DeleteById(db *sql.DB, id string) error {
	_, err := m.GetById(db, id)
	if err != nil {
		return err
	}

	query := "DELETE FROM motel_admins WHERE admin_id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = stmt.Query(id)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	return nil
}
