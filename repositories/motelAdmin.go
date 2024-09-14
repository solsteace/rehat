package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type MotelAdmin struct {
	Db *sql.DB
}

func (m MotelAdmin) Save(admin models.MotelAdmin) (int64, error) {
	query := `INSERT INTO motel_admins(user_id, motel_id) VALUES (?, ?)`
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	result, err := stmt.Exec(
		admin.UserID,
		admin.MotelID)
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, &ErrSQL{message: err.Error()}
	}

	return insertId, nil
}

func (m MotelAdmin) GetById(id int64) (models.MotelAdmin, error) {
	var motel models.MotelAdmin

	query := "SELECT * FROM motel_admins WHERE admin_id=?"
	stmt, err := m.Db.Prepare(query)
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
				Message: fmt.Sprintf("Couldn't find motel admin with id %d", id)}
		}
		return motel, &ErrSQL{message: err.Error()}
	}
	return motel, nil
}

func (m MotelAdmin) DeleteById(id int64) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := "DELETE FROM motel_admins WHERE admin_id=?"
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
