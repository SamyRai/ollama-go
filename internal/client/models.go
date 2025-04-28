package client

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// ListModels retrieves all available models.
func (c *OllamaClient) ListModels() (*structures.ModelListResponse, error) {
	c.Logger.Debug("Retrieving list of available models")

	var resp structures.ModelListResponse
	err := c.Request("GET", "/api/tags", nil, &resp)
	if err != nil {
		c.Logger.Error("Failed to list models: %v", err)
		return nil, err
	}

	c.Logger.Debug("Retrieved %d models", len(resp.Models))
	return &resp, nil
}

// ShowModel retrieves details about a specific model.
func (c *OllamaClient) ShowModel(req structures.ShowModelRequest) (*structures.ShowModelResponse, error) {
	c.Logger.Debug("Retrieving details for model: %s", req.Model)

	var resp structures.ShowModelResponse
	err := c.Request("POST", "/api/show", req, &resp)
	if err != nil {
		c.Logger.Error("Failed to get model details: %v", err)
		return nil, err
	}

	c.Logger.Debug("Successfully retrieved details for model: %s", req.Model)
	return &resp, nil
}
