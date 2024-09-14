package repositories

import "fmt"

type ErrSQL struct {
	message string
}

func (e *ErrSQL) Error() string {
	return fmt.Sprintf("[Uncategorized SQL Error] %s", e.message)
}

type ErrRecordNotFound struct {
	Message string
}

func (e *ErrRecordNotFound) Error() string {
	return e.Message
}

type ErrDuplicateEntry struct {
	Field string
}

func (e *ErrDuplicateEntry) Error() string {
	return fmt.Sprintf("Duplicate entry for `%s`", e.Field)
}
