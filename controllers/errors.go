package controllers

import "fmt"

type ErrAuth struct {
	message string
}

func (e *ErrAuth) Error() string {
	return fmt.Sprintf("Auth Error: %s", e.message)
}
