package errr

import (
	"errors"
	"fmt"
)

// StatusCodeError implement error interface
type StatusCodeError struct {
	StatusCode int
	Err        error
}

func (s StatusCodeError) Error() string {
	return s.Err.Error()
}

// New return StatusCodeError with same message, use this if we know exactly
// what error it is and what status code to return
func New(statusCode int, message string) StatusCodeError {
	return StatusCodeError{
		Err:        errors.New(message),
		StatusCode: statusCode,
	}
}

// Wrap transform error to StatusCodeError with same message,
// use this if we know exactly what status code to return
func Wrap(statusCode int, err error) StatusCodeError {
	return StatusCodeError{
		Err:        err,
		StatusCode: statusCode,
	}
}

func WrapF(statusCode int, format string, a ...any) StatusCodeError {
	return StatusCodeError{
		Err:        fmt.Errorf(format, a...),
		StatusCode: statusCode,
	}
}
