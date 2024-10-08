package repositories

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
	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.Motel{}.TableName())
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

func (m Motel) GetById(id int64) (models.Motel, error) {
	var motel models.Motel

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE motel_id=?",
		motel.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return motel, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&motel.MotelID,
		&motel.Name,
		&motel.Location,
		&motel.ContactNumber,
		&motel.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return motel, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find motel with id %d", id)}
		}
		return motel, &ErrSQL{message: err.Error()}
	}
	return motel, nil
}

func (m Motel) Save(motel models.Motel) (int64, error) {
	query := fmt.Sprintf(
		`INSERT 
			INTO %s(name, location, contact_number, email) 
			VALUES (?, ?, ?, ?)`,
		motel.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		motel.Name,
		motel.Location,
		motel.ContactNumber,
		motel.Email)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (m Motel) EditById(id int64, motel models.Motel) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		`UPDATE %s
			SET name = ?, 
				location = ?, 
				contact_number = ?, 
				email = ?,
				rating = ?
			WHERE motel_id = ?`,
		motel.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		motel.Name,
		motel.Location,
		motel.ContactNumber,
		motel.Email,
		motel.MotelID)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = result.RowsAffected()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	return nil
}

func (m Motel) DeleteById(id int64) error {
	motel, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE motel_id=?",
		motel.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Query(id)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	return nil
}
