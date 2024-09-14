package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type Room struct {
	Db *sql.DB
}

func (r Room) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	query := "SELECT * FROM rooms"
	rows, err := r.Db.Query(query)
	if err != nil {
		return rooms, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.RoomID,
			&room.ClassID,
			&room.RoomNumber,
			&room.Status)
		if err != nil {
			return rooms, &ErrSQL{message: err.Error()}
		}

		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return rooms, &ErrSQL{message: err.Error()}
	}

	return rooms, nil
}

func (r Room) GetById(id int64) (models.Room, error) {
	var room models.Room
	query := "SELECT * FROM rooms WHERE id=?"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return room, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&room.RoomID,
		&room.ClassID,
		&room.RoomNumber,
		&room.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return room, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find class with id %d", id)}
		}
		return room, &ErrSQL{message: err.Error()}
	}
	return room, nil
}

func (r Room) Save(room *models.Room) error {
	query := `INSERT INTO rooms(class_id, room_number, status) VALUES (?, ?, ?)`
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		room.ClassID,
		room.RoomNumber,
		room.Status)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	room.RoomID = insertId
	return nil
}

func (r Room) EditById(id int64, room *models.Room) error {
	query := `UPDATE rooms
				SET
					class_id = ?, 
					room_number = ?, 
					status = ?
				WHERE 
					room_id = ?`

	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		room.ClassID,
		room.RoomNumber,
		room.Status)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = result.RowsAffected()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	return nil
}

func (r Room) DeleteById(id int64) error {
	query := "DELETE FROM rooms WHERE id=?"
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
