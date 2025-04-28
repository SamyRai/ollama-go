package client

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// GetVersion retrieves the current API version.
func (c *OllamaClient) GetVersion() (*structures.VersionResponse, error) {
	c.Logger.Debug("Retrieving API version")

	var resp structures.VersionResponse
	err := c.Request("GET", "/api/version", nil, &resp)
	if err != nil {
		c.Logger.Error("Failed to get API version: %v", err)
		return nil, err
	}

	c.Logger.Debug("API version: %s", resp.Version)
	return &resp, nil
}

// GetRunningProcesses retrieves the list of running processes.
func (c *OllamaClient) GetRunningProcesses() (*structures.ModelProcessResponse, error) {
	c.Logger.Debug("Retrieving list of running processes")

	var resp structures.ModelProcessResponse
	err := c.Request("GET", "/api/ps", nil, &resp)
	if err != nil {
		c.Logger.Error("Failed to get running processes: %v", err)
		return nil, err
	}

	c.Logger.Debug("Retrieved %d running processes", len(resp.Models))
	return &resp, nil
}
