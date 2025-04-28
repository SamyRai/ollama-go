package api

import (
	"context"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/structures"
)

// ModelListResponse contains available local models.
type ModelListResponse = structures.ModelListResponse

// ModelInfo contains details about a local model.
type ModelInfo = structures.ModelInfo

// ShowModelResponse contains model details.
type ShowModelResponse = structures.ShowModelResponse

// ModelManager provides methods for model management.
type ModelManager struct {
	client *client.OllamaClient
}

// NewModelManager creates a new ModelManager.
func NewModelManager(client *client.OllamaClient) *ModelManager {
	return &ModelManager{
		client: client,
	}
}

// List retrieves all available models.
func (m *ModelManager) List(_ context.Context) (*ModelListResponse, error) {
	return m.client.ListModels()
}

// Show retrieves details about a specific model.
func (m *ModelManager) Show(ctx context.Context, modelName string) (*ShowModelResponse, error) {
	// Context could be used for cancellation or timeouts
	// This keeps the parameter for API consistency and future extensions
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	req := structures.ShowModelRequest{
		Model: modelName,
	}
	return m.client.ShowModel(req)
}

// Create creates a new model.
func (m *ModelManager) Create(ctx context.Context, name string, owner string) error {
	// Check for context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	req := structures.ModelManagementRequest{
		Name:  name,
		Owner: owner,
	}
	return m.client.CreateModel(req)
}

// Delete deletes an existing model.
func (m *ModelManager) Delete(ctx context.Context, modelName string) error {
	// Check for context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return m.client.DeleteModel(modelName)
}

// Copy copies a model to a new name.
func (m *ModelManager) Copy(ctx context.Context, sourceModel, targetModel string) error {
	// Check for context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return m.client.CopyModel(sourceModel, targetModel)
}

// Pull pulls a model from a remote repository.
func (m *ModelManager) Pull(ctx context.Context, modelName string) error {
	// Check for context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return m.client.PullModel(modelName)
}

// Push pushes a model to a remote repository.
func (m *ModelManager) Push(ctx context.Context, modelName string) error {
	// Check for context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return m.client.PushModel(modelName)
}
