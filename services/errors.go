package services

import "fmt"

type ErrSQL struct {
	message string
}

func (e *ErrSQL) Error() string {
	return fmt.Sprintf("SQL Error: %s", e.message)
}

type ErrServiceNotImplemented struct{}

func (e *ErrServiceNotImplemented) Error() string {
	return "Service is not implemented yet"
}
