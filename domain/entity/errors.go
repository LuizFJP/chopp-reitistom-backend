package entity

import (
	"errors"
	"strings"

	"github.com/vardius/trace"
)

// Application errors
var (
	ErrInvalid           = errors.New("validation failed")
	ErrUnauthorized      = errors.New("access denied")
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrInternal          = errors.New("internal system error")
	ErrTemporaryDisabled = errors.New("temporary disabled")
	ErrTimeout           = errors.New("timeout")
	ErrBodyIsMissing     = errors.New("request body is missing")
	ErrUUIDIsMissing     = errors.New("uuid is missing")
)

// New returns new app error that formats as the given text.
func New(message string) *AppError {
	return newAppError(errors.New(message))
}

// Wrap returns new app error wrapping target error.
// If passed value is nil will fallback to internal
func Wrap(err error) *AppError {
	return newAppError(err)
}

func newAppError(err error) *AppError {
	if err == nil {
		err = ErrInternal
	}

	return &AppError{
		err:   err,
		trace: trace.FromParent(2, trace.Lfile|trace.Lline),
	}
}

type AppError struct {
	trace string
	err   error
}

func (e *AppError) Error() string {
	return e.err.Error()
}

func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) StackTrace() string {
	var stack []string

	if e.trace != "" {
		stack = append(stack, e.trace)
	}

	if e.err != nil {
		var next *AppError
		if errors.As(e.err, &next) {
			stack = append(stack, next.StackTrace())
		}
	}

	return strings.Join(stack, "\n")
}
