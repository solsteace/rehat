package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type Class struct {
	Db *sql.DB
}

func (r Class) GetAll() ([]models.RoomClass, error) {
	var classes []models.RoomClass
	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.RoomClass{}.TableName())
	rows, err := r.Db.Query(query)
	if err != nil {
		return classes, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var class models.RoomClass
		err := rows.Scan(
			&class.MotelID,
			&class.Name,
			&class.Price)
		if err != nil {
			return classes, &ErrSQL{message: err.Error()}
		}

		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return classes, &ErrSQL{message: err.Error()}
	}

	return classes, nil
}

func (r Class) GetById(id int64) (models.RoomClass, error) {
	var class models.RoomClass
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE id=?",
		class.TableName())
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return class, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&class.MotelID,
		&class.Name,
		&class.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return class, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find class with id %d", id)}
		}
		return class, &ErrSQL{message: err.Error()}
	}
	return class, nil
}

func (r Class) Save(class *models.RoomClass) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(motel_id, name, price) VALUES (?, ?, ?)`,
		class.TableName())
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		class.MotelID,
		class.Name,
		class.Price)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	class.ClassID = insertId
	return nil
}

func (r Class) EditById(id int64, class *models.RoomClass) error {
	query := fmt.Sprintf(
		`UPDATE %s
			SET motel_id = ?, 
				name = ?, 
				price = ?
			WHERE class_id = ?`,
		class.TableName())
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		class.MotelID,
		class.Name,
		class.Price)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = result.RowsAffected()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	return nil
}

func (r Class) DeleteById(id int64) error {
	class, err := r.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id=?",
		class.TableName())
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = result.RowsAffected()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	return nil
}
