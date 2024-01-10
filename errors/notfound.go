package errors

// NotFoundError is a ProcessOut resource not found error
type NotFoundError struct {
	message string
	code    string
}

// Error returns the error message
func (e *NotFoundError) Error() string {
	return e.message
}

// Code returns the error code returned by ProcessOut
func (e *NotFoundError) Code() string {
	return e.code
}
