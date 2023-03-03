package apistatus

import (
	"errors"
	"fmt"
)

/*
* Returned models error to client.
 */

// NewErrorAPI return error to client in proper manner.
func NewErrorAPI(err error) *ErrorAPI {
	return &ErrorAPI{
		Err: err,
	}
}

// NewInvalidArgsAPI return error to client in good manner.
func NewInvalidArgsAPI(inv []InvalidArgument, err *Error) *ErrorInvalidArgs {
	return &ErrorInvalidArgs{
		Error:   err,
		InvArgs: inv,
	}
}

/*
* Error "Factories"
 */

// NewAuthorization to create a 401.
func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

// NewBadRequest to create 400 errors (validation, for example).
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflict to create an error for 409.
func NewConflict(name, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewInternal for 500 errors and unknown errors.
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

// NewNotFound to create an error for 404.
func NewNotFound(name, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

// IsNotFound return boolean and compare error with IsNotFoundError
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	var e *Error
	if ok := errors.As(err, &e); !ok {
		return false
	}

	return e.Type == NotFound
}

// NewPayloadTooLarge to create an error for 413.
func NewPayloadTooLarge(maxBodySize, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}
}

// NewServiceUnavailable to create an error for 503.
func NewServiceUnavailable() *Error {
	return &Error{
		Type:    ServiceUnavailable,
		Message: "Service unavailable or timed out",
	}
}

// NewUnsupportedMediaType to create an error for 415.
func NewUnsupportedMediaType(reason string) *Error {
	return &Error{
		Type:    UnsupportedMediaType,
		Message: reason,
	}
}

// NewSuccessStatus return status to client in proper manner
func NewSuccessStatus(msg string) *SuccessStatus {
	return &SuccessStatus{
		Status: msg,
	}
}
