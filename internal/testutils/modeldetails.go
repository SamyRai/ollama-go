package testutils

import "github.com/SamyRai/ollama-go/internal/structures"

// ModelDetails is a placeholder for the ModelDetails struct used in the mock server.
// This is needed because the actual ModelDetails struct is not exported from the structures package.
type ModelDetails struct{}

// NewModelDetails creates a new ModelDetails struct.
func NewModelDetails() structures.ModelDetails {
	return structures.ModelDetails{}
}
