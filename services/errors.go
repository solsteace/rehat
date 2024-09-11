package services

import "fmt"

type ErrSQL struct {
	message string
}

func (e *ErrSQL) Error() string {
	return fmt.Sprintf("[Uncategorized SQL Error] %s", e.message)
}

type ErrNotImplemented struct{}

func (e *ErrNotImplemented) Error() string {
	return "Service is not implemented yet"
}
