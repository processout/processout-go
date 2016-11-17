package errors

// AuthenticationError is a ProcessOut authentication error
type AuthenticationError struct {
	message string
	code    string
}

// Error returns the error message
func (e *AuthenticationError) Error() string {
	return e.message
}
