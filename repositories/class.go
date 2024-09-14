package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type Class struct {
	Db *sql.DB
}

func (r Class) GetAll() ([]models.Class, error) {
	var classes []models.Class
	query := "SELECT * FROM classes"
	rows, err := r.Db.Query(query)
	if err != nil {
		return classes, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var class models.Class
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

func (r Class) GetById(id int64) (models.Class, error) {
	var class models.Class
	query := "SELECT * FROM classes WHERE id=?"
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

func (r Class) Save(class *models.Class) error {
	query := `INSERT INTO classes(motel_id, name, price) VALUES (?, ?, ?)`
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

func (r Class) EditById(id int64, class *models.Class) error {
	query := `UPDATE classes
				SET
					motel_id = ?, 
					name = ?, 
					price = ?
				WHERE 
					class_id = ?`

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
	query := "DELETE FROM classes WHERE id=?"
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
