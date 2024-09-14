package models

import "time"

type Reservation struct {
	ReservationID int64     `json:"reservation_id" db:"reservation_id"`
	RoomID        int64     `json:"room_id" db:"room_id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	ReserveStart  time.Time `json:"reserve_start" db:"reserve_start"`
	ReserveEnd    time.Time `json:"reserve_end" db:"reserve_end"`
	Checkout      time.Time `json:"checkout" db:"checkout"`
	Total         int       `json:"total" db:"total"`
}
