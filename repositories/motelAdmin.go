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
	defer stmt.Close()

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
	var admin models.MotelAdmin

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE admin_id=?",
		admin.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return admin, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&admin.AdminID,
		&admin.UserID,
		&admin.MotelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return admin, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find motel admin with id %d", id)}
		}
		return admin, &ErrSQL{message: err.Error()}
	}
	return admin, nil
}

func (m MotelAdmin) GetByUserAndMotelId(userId, motelId int64) (models.MotelAdmin, error) {
	var admin models.MotelAdmin

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id=? AND motel_id=?",
		admin.TableName())
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return admin, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId, motelId).Scan(
		&admin.AdminID,
		&admin.UserID,
		&admin.MotelID)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return admin, &ErrRecordNotFound{
				Message: fmt.Sprintf(
					"Couldn't find motel admin with admin id %d and motel id %d",
					userId, motelId)}
		}
		return admin, &ErrSQL{message: err.Error()}
	}
	return admin, nil
}

func (m MotelAdmin) DeleteById(id int64) error {
	admin, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE admin_id=?",
		admin.TableName())
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

func (m MotelAdmin) DeleteByMotelId(id int64) error {
	admin, err := m.GetById(id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE motel_id=?",
		admin.TableName())
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
