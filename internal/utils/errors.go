package utils

import (
	"errors"
	"fmt"
	"net/http"
)

// Common error types that can be used for type assertions
var (
	ErrInvalidResponse = errors.New("received an invalid response from the API")
	ErrRequestFailed   = errors.New("API request failed")
	ErrTimeout         = errors.New("request timed out")
	ErrModelNotFound   = errors.New("specified model was not found")
	ErrInvalidArgument = errors.New("invalid argument provided")
	ErrServerError     = errors.New("server error occurred")
	ErrNetworkError    = errors.New("network error occurred")
)

// APIError represents an error returned by the Ollama API
type APIError struct {
	StatusCode int
	Status     string
	Message    string
	Err        error
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error: %s (status: %s, code: %d)", e.Message, e.Status, e.StatusCode)
	}
	return fmt.Sprintf("API error: %s (code: %d)", e.Status, e.StatusCode)
}

// Unwrap returns the wrapped error
func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new API error from an HTTP response
func NewAPIError(resp *http.Response, message string) *APIError {
	return &APIError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Message:    message,
		Err:        ErrRequestFailed,
	}
}

// ValidationError represents an error due to invalid input
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s (field: %s, value: %v)", e.Message, e.Field, e.Value)
}

// Unwrap returns the wrapped error
func (e *ValidationError) Unwrap() error {
	return ErrInvalidArgument
}

// NewValidationError creates a new validation error
func NewValidationError(field string, value interface{}, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}
