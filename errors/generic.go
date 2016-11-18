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

// Code returns the error code returned by ProcessOut
func (e *Error) Code() string {
	return e.code
}

// Root returns the root error, if any
func (e *Error) Root() error {
	return e.err
}
