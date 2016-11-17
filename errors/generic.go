package errors

// Error is a generic ProcessOut error
type Error struct {
	err     error
	message string
	code    string
}

// Error returns the error message
func (e *Error) Error() string {
	return e.message
}
