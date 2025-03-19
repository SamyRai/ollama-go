package client

import (
	"hrelay/core/llm/ollama/structures"
)

// ListModels retrieves all available models.
func (c *OllamaClient) ListModels() (*structures.ModelListResponse, error) {
	var resp structures.ModelListResponse
	err := c.Request("GET", "/api/tags", nil, &resp)
	return &resp, err
}

// ShowModel retrieves details about a specific model.
func (c *OllamaClient) ShowModel(req structures.ShowModelRequest) (*structures.ShowModelResponse, error) {
	var resp structures.ShowModelResponse
	err := c.Request("POST", "/api/show", req, &resp)
	return &resp, err
}
