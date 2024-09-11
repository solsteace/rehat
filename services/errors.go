package services

import "fmt"

type ErrSQL struct {
	message string
}

func (e *ErrSQL) Error() string {
	return fmt.Sprintf("[Uncategorized SQL Error] %s", e.message)
}

type ErrRecordNotFound struct {
	message string
}

func (e *ErrRecordNotFound) Error() string {
	return e.message
}

type ErrNotImplemented struct{}

func (e *ErrNotImplemented) Error() string {
	return "Service is not implemented yet"
}

type ErrDuplicateEntry struct {
	field string
}

func (e *ErrDuplicateEntry) Error() string {
	return fmt.Sprintf("Duplicate entry for `%s`", e.field)
}
