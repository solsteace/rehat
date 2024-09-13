package models

import (
	"database/sql"
	"fmt"
)

type Motel struct {
	MotelID       int64  `json:"motel_id" db:"motel_id"`
	Name          string `json:"name" db:"name"`
	Location      string `json:"location" db:"location"`
	ContactNumber string `json:"contact_number" db:"contact_number"`
	Email         string `json:"email" db:"email"`
}

func (m Motel) GetAll(db *sql.DB) ([]Motel, error) {
	var motels []Motel
	query := "SELECT * FROM motels"
	rows, err := db.Query(query)
	if err != nil {
		return motels, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var motel Motel
		err := rows.Scan(
			&motel.MotelID,
			&motel.Name,
			&motel.Location,
			&motel.ContactNumber,
			&motel.Email)
		if err != nil {
			return motels, &ErrSQL{message: err.Error()}
		}

		motels = append(motels, motel)
	}

	if err := rows.Err(); err != nil {
		return motels, &ErrSQL{message: err.Error()}
	}
	return motels, nil
}

func (m Motel) Save(db *sql.DB) (int64, error) {
	query := `INSERT INTO motels(name, location, contact_number, email) 
				VALUES (?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		m.Name,
		m.Location,
		m.ContactNumber,
		m.Email)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (m Motel) GetById(db *sql.DB, id string) (Motel, error) {
	var motel Motel

	query := "SELECT * FROM motels WHERE motel_id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return motel, &ErrSQL{message: err.Error()}
	}

	err = stmt.QueryRow(id).Scan(
		&motel.MotelID,
		&motel.Name,
		&motel.Location,
		&motel.ContactNumber,
		&motel.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return motel, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find motel with id %s", id)}
		}
		return motel, &ErrSQL{message: err.Error()}
	}
	return motel, nil
}

func (m Motel) EditById(db *sql.DB, id string) (int64, error) {
	_, err := m.GetById(db, id)
	if err != nil {
		return 0, err
	}

	query := `UPDATE motels
				SET
					name = ?, 
					location = ?, 
					contact_number = ?, 
					email = ?
				WHERE 
					motel_id = ?  `
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		m.Name,
		m.Location,
		m.ContactNumber,
		m.Email,
		m.MotelID)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return affectedRows, nil
}

func (m Motel) DeleteById(db *sql.DB, id string) error {
	_, err := m.GetById(db, id)
	if err != nil {
		return err
	}

	query := "DELETE FROM motels WHERE motel_id=?"
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
