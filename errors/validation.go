package errors

// ValidationError is a ProcessOut validation error
type ValidationError struct {
	message string
	code    string
}

// Error returns the error message
func (e *ValidationError) Error() string {
	return e.message
}
