package errors

// InternalError is a ProcessOut internal error
type InternalError struct {
	message string
	code    string
}

// Error returns the error message
func (e *InternalError) Error() string {
	return e.message
}

// Code returns the error code returned by ProcessOut
func (e *InternalError) Code() string {
	return e.code
}
