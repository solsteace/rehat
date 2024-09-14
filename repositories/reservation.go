package repositories

import (
	"database/sql"
	"fmt"

	"github.com/solsteace/rest/models"
)

type Reservation struct {
	Db *sql.DB
}

func (r Reservation) GetAll() ([]models.Reservation, error) {
	var reservations []models.Reservation
	query := "SELECT * FROM reservations"
	rows, err := r.Db.Query(query)
	if err != nil {
		return reservations, &ErrSQL{message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(
			&reservation.RoomID,
			&reservation.UserID,
			&reservation.ReserveStart,
			&reservation.ReserveEnd,
			&reservation.Checkout,
			&reservation.Total)
		if err != nil {
			return reservations, &ErrSQL{message: err.Error()}
		}

		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return reservations, &ErrSQL{message: err.Error()}
	}

	return reservations, nil
}

func (r Reservation) GetById(id int64) (models.Reservation, error) {
	var reservation models.Reservation
	query := "SELECT * FROM reservations WHERE id=?"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return reservation, &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&reservation.RoomID,
		&reservation.UserID,
		&reservation.ReserveStart,
		&reservation.ReserveEnd,
		&reservation.Checkout,
		&reservation.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return reservation, &ErrRecordNotFound{
				Message: fmt.Sprintf("Couldn't find reservation with id %d", id)}
		}
		return reservation, &ErrSQL{message: err.Error()}
	}
	return reservation, nil
}

func (r Reservation) Save(reservation *models.Reservation) error {
	query := `INSERT 
				INTO reservations( room_id, user_id, reserve_start, 
					reserve_end, checkout, total) 
				VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		reservation.RoomID,
		reservation.UserID,
		reservation.ReserveStart,
		reservation.ReserveEnd,
		reservation.Checkout,
		reservation.Total)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	reservation.ReservationID = insertId
	return nil
}

func (r Reservation) EditById(id int64, reservation *models.Reservation) error {
	query := `UPDATE reservations
				SET 
					room_id = ?, 
					user_id = ?, 
					reserve_start = ?, 
					reserve_end = ?, 
					checkout = ?, 
					total = ?
				WHERE 
					reservation_id = ?`

	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		reservation.RoomID,
		reservation.UserID,
		reservation.ReserveStart,
		reservation.ReserveEnd,
		reservation.Checkout,
		reservation.Total)
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	_, err = result.RowsAffected()
	if err != nil {
		return &ErrSQL{message: err.Error()}
	}

	return nil
}

func (r Reservation) DeleteById(id int64) error {
	query := "DELETE FROM reservations WHERE id=?"
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
