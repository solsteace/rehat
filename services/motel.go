package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/solsteace/rest/models"
)

const TABLE_NAME = "motels"

type Motel struct {
	Db *sql.DB
}

func (r Motel) GetAll() ([]models.Motel, error) {
	var motels []models.Motel

	query := fmt.Sprintf("SELECT * FROM %s", TABLE_NAME)
	rows, err := r.Db.Query(query)
	if err != nil {
		return motels, err
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
			return motels, err
		}

		motels = append(motels, motel)
	}

	if err := rows.Err(); err != nil {
		return motels, err
	}
	return motels, nil
}

// TODO: Complete
func (r Motel) Create() ([]models.Motel, error) {
	var motel []models.Motel
	return motel, errors.New("NOT IMPLEMENTED")
}

func (r Motel) GetById(id string) (models.Motel, error) {
	var motel models.Motel

	query := fmt.Sprintf("SELECT * FROM %s WHERE motel_id=?", TABLE_NAME)
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return motel, nil
	}

	err = stmt.QueryRow(id).Scan(
		&motel.MotelID,
		&motel.Name,
		&motel.Location,
		&motel.ContactNumber,
		&motel.Email)
	if err != nil {
		return motel, err
	}
	return motel, nil
}

// TODO: Complete
func (r Motel) EditById(id string) (models.Motel, error) {
	var motel models.Motel
	return motel, errors.New("NOT IMPLEMENTED")
}

func (r Motel) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE motel_id=?", TABLE_NAME)
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return nil
	}

	_, err = stmt.Query(id)
	if err != nil {
		return err
	}
	return nil
}
