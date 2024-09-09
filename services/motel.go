package services

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type Motel struct {
	Db *sql.DB
}

func (m Motel) GetAll() ([]models.Motel, error) {
	var motels []models.Motel
	query := "SELECT * FROM motels"
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
	query := `INSERT INTO motels(name, location, contact_number, email) 
				VALUES (?, ?, ?, ?)`
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

	query := "SELECT * FROM motels WHERE motel_id=?"
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
				message: fmt.Sprintf("Couldn't find motel with id %s", id)}
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

	query := `UPDATE motels
				SET
					name = ?, 
					location = ?, 
					contact_number = ?, 
					email = ?
				WHERE 
					motel_id = ?  `
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

	query := "DELETE FROM motels WHERE motel_id=?"
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
