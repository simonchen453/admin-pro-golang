package errors

import (
	"fmt"
)

// Common error definitions
var (
	// Authentication errors
	ErrUserNotFound    = fmt.Errorf("user not found")
	ErrInvalidPassword = fmt.Errorf("invalid password")
	ErrInvalidToken    = fmt.Errorf("invalid token")
	ErrTokenExpired    = fmt.Errorf("token expired")
	ErrUnauthorized    = fmt.Errorf("unauthorized access")

	// Validation errors
	ErrInvalidInput  = fmt.Errorf("invalid input")
	ErrInvalidParams = fmt.Errorf("invalid parameters")
	ErrRequiredField = fmt.Errorf("required field missing")

	// Resource errors
	ErrNotFound      = fmt.Errorf("resource not found")
	ErrAlreadyExists = fmt.Errorf("resource already exists")
	ErrConflict      = fmt.Errorf("resource conflict")

	// Business logic errors
	ErrInvalidOperation = fmt.Errorf("invalid operation")
	ErrPermissionDenied = fmt.Errorf("permission denied")
	ErrStatusInvalid    = fmt.Errorf("invalid status")

	// Database errors
	ErrDatabase       = fmt.Errorf("database error")
	ErrRecordNotFound = fmt.Errorf("record not found")
	ErrDuplicateEntry = fmt.Errorf("duplicate entry")
)

// Wrap wraps an error with additional context
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// Wrapf wraps an error with formatted message
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
