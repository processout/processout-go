package errors

// CodedError is the interface implemented by ProcessOut errors
type CodedError interface {
	Error() string
	Code() string
}

// New creates a new ProcessOut error from an error
func New(err error, code, message string) error {
	if err != nil {
		message = err.Error()
	}
	return &Error{
		err:     err,
		message: message,
		code:    code,
	}
}

// NewFromResponse creates an error from a response data
func NewFromResponse(status int, code, message string) error {
	if status == 404 {
		return &NotFoundError{
			message: message,
			code:    code,
		}
	}
	if status == 401 {
		return &AuthenticationError{
			message: message,
			code:    code,
		}
	}
	if status == 400 {
		return &ValidationError{
			message: message,
			code:    code,
		}
	}
	if status >= 500 {
		return &InternalError{
			message: message,
			code:    code,
		}
	}

	return New(nil, code, message)
}
