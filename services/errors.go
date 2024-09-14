package services

type ErrNotImplemented struct{}

func (e *ErrNotImplemented) Error() string {
	return "Service is not implemented yet"
}
