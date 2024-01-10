package errors

// NetworkError is a ProcessOut network error
type NetworkError struct {
	err     error
	message string
	code    string
}

// Error returns the error message
func (e *NetworkError) Error() string {
	return e.message
}

// Code returns the error code returned by ProcessOut
func (e *NetworkError) Code() string {
	return e.code
}
