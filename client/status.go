package client

import (
	"github.com/SamyRai/ollama-go/structures"
)

// GetVersion retrieves the current API version.
func (c *OllamaClient) GetVersion() (*structures.VersionResponse, error) {
	var resp structures.VersionResponse
	err := c.Request("GET", "/api/version", nil, &resp)
	return &resp, err
}

// GetRunningProcesses retrieves the list of running structures.
func (c *OllamaClient) GetRunningProcesses() (*structures.ModelProcessResponse, error) {
	var resp structures.ModelProcessResponse
	err := c.Request("GET", "/api/ps", nil, &resp)
	return &resp, err
}
