package models

type Motel struct {
	MotelID       int64  `json:"motel_id" db:"motel_id"`
	Name          string `json:"name" db:"name"`
	Location      string `json:"location" db:"location"`
	ContactNumber string `json:"contact_number" db:"contact_number"`
	Email         string `json:"email" db:"email"`
	Rating        int    `json:"rating" db:"rating"`
}
