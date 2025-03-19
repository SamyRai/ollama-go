package client

import (
	"hrelay/core/llm/ollama/structures"
)

// CreateModel sends a request to create a new model.
func (c *OllamaClient) CreateModel(req structures.ModelManagementRequest) error {
	return c.Request("POST", "/api/create", req, nil)
}

// DeleteModel sends a request to delete an existing model.
func (c *OllamaClient) DeleteModel(modelName string) error {
	payload := map[string]string{"model": modelName}
	return c.Request("DELETE", "/api/delete", payload, nil)
}

// CopyModel copies a model to a new name.
func (c *OllamaClient) CopyModel(sourceModel, targetModel string) error {
	payload := map[string]string{"sourceModel": sourceModel, "targetModel": targetModel}
	return c.Request("POST", "/api/copy", payload, nil)
}

// PullModel pulls a model from a remote repository.
func (c *OllamaClient) PullModel(modelName string) error {
	payload := map[string]string{"model": modelName}
	return c.Request("POST", "/api/pull", payload, nil)
}

// PushModel pushes a model to a remote repository.
func (c *OllamaClient) PushModel(modelName string) error {
	payload := map[string]string{"model": modelName}
	return c.Request("POST", "/api/push", payload, nil)
}
