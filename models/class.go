package models

type Class struct {
	ClassID int64  `json:"class_id" db:"class_id"`
	MotelID int64  `json:"motel_id" db:"motel_id"`
	Name    string `json:"class_name" db:"class_name"`
	Price   int    `json:"price" db:"price"`
}
