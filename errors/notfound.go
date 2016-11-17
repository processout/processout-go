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
