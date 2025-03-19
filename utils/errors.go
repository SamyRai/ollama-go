package utils

import "errors"

// Predefined errors for the Ollama client.
var (
    ErrInvalidResponse = errors.New("received an invalid response from the API")
    ErrRequestFailed   = errors.New("API request failed")
    ErrTimeout         = errors.New("request timed out")
    ErrModelNotFound   = errors.New("specified model was not found")
)
