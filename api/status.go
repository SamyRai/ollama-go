package api

import (
	"context"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/structures"
)

// VersionResponse returns the server version.
type VersionResponse = structures.VersionResponse

// ModelProcessResponse lists running models.
type ModelProcessResponse = structures.ModelProcessResponse

// ModelProcess contains details about a running model.
type ModelProcess = structures.ModelProcess

// StatusManager provides methods for status-related operations.
type StatusManager struct {
	client *client.OllamaClient
}

// NewStatusManager creates a new StatusManager.
func NewStatusManager(client *client.OllamaClient) *StatusManager {
	return &StatusManager{
		client: client,
	}
}

// GetVersion retrieves the current API version.
func (s *StatusManager) GetVersion(_ context.Context) (*VersionResponse, error) {
	return s.client.GetVersion()
}

// GetRunningProcesses retrieves the list of running models.
func (s *StatusManager) GetRunningProcesses(_ context.Context) (*ModelProcessResponse, error) {
	return s.client.GetRunningProcesses()
}
