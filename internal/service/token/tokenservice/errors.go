package tokenservice

import (
	"errors"
)

// InvalidTokenError is used to indicate an error with a provided token
type InvalidTokenError struct {
	label string
}

// NewInvalidToken return invalid token error to client in proper manner
func NewInvalidToken(label string) *InvalidTokenError {
	return &InvalidTokenError{
		label: label,
	}
}

// IsInvalidToken return boolean and compare error with InvalidTokenError
func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}

	var e *InvalidTokenError

	return errors.As(err, &e)
}

// Error implements the error interface.
func (e *InvalidTokenError) Error() string {
	return "api: " + e.label + " token is invalid"
}
