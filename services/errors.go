package services

type ErrNotImplemented struct{}

func (e *ErrNotImplemented) Error() string {
	return "Service is not implemented yet"
}

type ErrNoResourcePermission struct{}

func (e *ErrNoResourcePermission) Error() string {
	return "User has no permission to the resource"
}

type ErrIllegalOperation struct {
	reason string
}

func (e *ErrIllegalOperation) Error() string {
	return e.reason
}
