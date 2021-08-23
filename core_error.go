package corerr

import (
	"fmt"
)

// CoreError ...
type CoreError struct {
	ErrorCode    int
	ErrorMessage string
	InnerError   error
}

// NewCoreError ...
func NewCoreError(code int, msg string) error {
	return &CoreError{
		ErrorCode:    code,
		ErrorMessage: msg,
	}
}

// NewWrappedCoreError ...
func NewWrappedCoreError(code int, msg string, err error) error {
	return &CoreError{
		ErrorCode:    code,
		ErrorMessage: msg,
		InnerError:   err,
	}
}

// NewIfNotCoreError will create an error only if err is not a Gauntlet error
func NewIfNotCoreError(code int, msg string, err error) error {
	if _, ok := err.(*CoreError); ok {
		return err
	}
	return NewWrappedCoreError(code, msg, err)
}

// Error ...
func (e *CoreError) Error() string {
	return fmt.Sprintf("%d: %s", e.ErrorCode, e.ErrorMessage)
}

// Cause ...
func (e *CoreError) Cause(err error) error {
	if e.InnerError != nil {
		return e.InnerError
	}
	return e
}

// Is ...
func (e *CoreError) Is(target error) bool {
	t, ok := target.(*CoreError)
	if !ok {
		return false
	}
	return t.ErrorCode == e.ErrorCode
}

// Unwrap ...
func (e *CoreError) Unwrap() error { return e.InnerError }

// HTTPStatusCode ...
func (e *CoreError) HTTPStatusCode() int {
	if e.ErrorCode < 20000 {
		return e.ErrorCode % 1000
	}
	return 500
}
