package client

import (
	"github.com/SamyRai/ollama-go/structures"
)

// CreateModel sends a request to create a new model via the /api/create endpoint.
func (c *OllamaClient) CreateModel(req structures.ModelManagementRequest) error {
	return c.Request("POST", "/api/create", req, nil)
}

// DeleteModel sends a request to delete an existing model via the /api/delete endpoint.
func (c *OllamaClient) DeleteModel(modelName string) error {
	payload := map[string]string{"name": modelName}
	return c.Request("DELETE", "/api/delete", payload, nil)
}

// CopyModel copies a model to a new name via the /api/copy endpoint.
func (c *OllamaClient) CopyModel(source, destination string) error {
	payload := map[string]string{"source": source, "destination": destination}
	return c.Request("POST", "/api/copy", payload, nil)
}

// PullModel pulls a model from a remote repository via the /api/pull endpoint.
func (c *OllamaClient) PullModel(modelName string) error {
	payload := map[string]string{"name": modelName}
	return c.Request("POST", "/api/pull", payload, nil)
}

// PushModel pushes a model to a remote repository via the /api/push endpoint.
func (c *OllamaClient) PushModel(modelName string) error {
	payload := map[string]string{"name": modelName}
	return c.Request("POST", "/api/push", payload, nil)
}
