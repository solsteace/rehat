package models

type Room struct {
	RoomID     int64  `json:"room_id" db:"room_id"`
	ClassID    int64  `json:"class_id" db:"class_id"`
	RoomNumber int    `json:"room_number" db:"room_number"`
	Status     string `json:"status" db:"status"`
}
