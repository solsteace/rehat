package services

import "fmt"

type ErrSQL struct {
	message string
}

func (e *ErrSQL) Error() string {
	return fmt.Sprintf("[Uncategorized SQL Error] %s", e.message)
}

type ErrSQLNoRows struct {
	message string
}

func (e *ErrSQLNoRows) Error() string {
	return e.message
}

type ErrServiceNotImplemented struct{}

func (e *ErrServiceNotImplemented) Error() string {
	return "Service is not implemented yet"
}
