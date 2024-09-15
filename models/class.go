package models

type RoomClass struct {
	ClassID int64  `json:"class_id" db:"class_id"`
	MotelID int64  `json:"motel_id" db:"motel_id"`
	Name    string `json:"name" db:"name"`
	Price   int    `json:"price" db:"price"`
}

func (rc RoomClass) TableName() string { return "room_classes" }
