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

// Code returns the error code returned by ProcessOut
func (e *AuthenticationError) Code() string {
	return e.code
}
