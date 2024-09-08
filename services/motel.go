package services

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

const TABLE_NAME = "motels"

type Motel struct {
	Db *sql.DB
}

func (m Motel) GetAll() ([]models.Motel, error) {
	var motels []models.Motel
	query := fmt.Sprintf("SELECT * FROM %s", TABLE_NAME)
	rows, err := m.Db.Query(query)
	if err != nil {
		return motels, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var motel models.Motel
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

func (m Motel) Create(newMotel models.Motel) (int64, error) {
	query := fmt.Sprintf(
		`INSERT INTO %s(name, location, contact_number, email) 
			VALUES (?, ?, ?, ?)
		`, TABLE_NAME)
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		newMotel.Name,
		newMotel.Location,
		newMotel.ContactNumber,
		newMotel.Email)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (m Motel) GetById(id string) (models.Motel, error) {
	var motel models.Motel

	query := fmt.Sprintf("SELECT * FROM %s WHERE motel_id=?", TABLE_NAME)
	stmt, err := m.Db.Prepare(query)
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
			return motel, &ErrSQLNoRows{
				message: fmt.Sprintf(
					"Couldn't find %s with id %s", TABLE_NAME, id)}
		}
		return motel, &ErrSQL{message: err.Error()}
	}
	return motel, nil
}

func (m Motel) EditById(id string, newMotel models.Motel) (int64, error) {
	_, err := m.GetById(id)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(
		`UPDATE %s 
			SET
				name = ?, 
				location = ?, 
				contact_number = ?, 
				email = ?
			WHERE 
				motel_id = ?
		`, TABLE_NAME)
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		newMotel.Name,
		newMotel.Location,
		newMotel.ContactNumber,
		newMotel.Email,
		newMotel.MotelID)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return affectedRows, nil
}

func (m Motel) DeleteById(id string) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE motel_id=?", TABLE_NAME)
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = stmt.Query(id)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	return nil
}
