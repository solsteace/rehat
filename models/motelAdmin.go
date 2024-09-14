package models

type MotelAdmin struct {
	AdminID int64 `json:"admin_id" db:"admin_id"`
	UserID  int64 `json:"user_id" db:"user_id"`
	MotelID int64 `json:"motel_id" db:"motel_id"`
}
