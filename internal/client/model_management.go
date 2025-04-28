package client

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// CreateModel sends a request to create a new model.
func (c *OllamaClient) CreateModel(req structures.ModelManagementRequest) error {
	c.Logger.Debug("Creating new model: %s", req.Name)

	err := c.Request("POST", "/api/create", req, nil)
	if err != nil {
		c.Logger.Error("Failed to create model: %v", err)
		return err
	}

	c.Logger.Info("Successfully created model: %s", req.Name)
	return nil
}

// DeleteModel sends a request to delete an existing model.
func (c *OllamaClient) DeleteModel(modelName string) error {
	c.Logger.Debug("Deleting model: %s", modelName)

	payload := map[string]string{"model": modelName}
	err := c.Request("DELETE", "/api/delete", payload, nil)
	if err != nil {
		c.Logger.Error("Failed to delete model: %v", err)
		return err
	}

	c.Logger.Info("Successfully deleted model: %s", modelName)
	return nil
}

// CopyModel copies a model to a new name.
func (c *OllamaClient) CopyModel(sourceModel, targetModel string) error {
	c.Logger.Debug("Copying model from %s to %s", sourceModel, targetModel)

	payload := map[string]string{"sourceModel": sourceModel, "targetModel": targetModel}
	err := c.Request("POST", "/api/copy", payload, nil)
	if err != nil {
		c.Logger.Error("Failed to copy model: %v", err)
		return err
	}

	c.Logger.Info("Successfully copied model from %s to %s", sourceModel, targetModel)
	return nil
}

// PullModel pulls a model from a remote repository.
func (c *OllamaClient) PullModel(modelName string) error {
	c.Logger.Debug("Pulling model: %s", modelName)

	payload := map[string]string{"model": modelName}
	err := c.Request("POST", "/api/pull", payload, nil)
	if err != nil {
		c.Logger.Error("Failed to pull model: %v", err)
		return err
	}

	c.Logger.Info("Successfully pulled model: %s", modelName)
	return nil
}

// PushModel pushes a model to a remote repository.
func (c *OllamaClient) PushModel(modelName string) error {
	c.Logger.Debug("Pushing model: %s", modelName)

	payload := map[string]string{"model": modelName}
	err := c.Request("POST", "/api/push", payload, nil)
	if err != nil {
		c.Logger.Error("Failed to push model: %v", err)
		return err
	}

	c.Logger.Info("Successfully pushed model: %s", modelName)
	return nil
}
