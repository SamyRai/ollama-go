package tests

import (
	"context"
	"testing"

	"github.com/SamyRai/ollama-go/api"
	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelManager(t *testing.T) {
	// Create a context
	ctx := context.Background()

	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create the model manager
	manager := api.NewModelManager(cli)

	// Test cases for model operations
	tests := []struct {
		name        string
		operation   func(ctx context.Context) (interface{}, error)
		validate    func(t *testing.T, result interface{})
		expectError bool
	}{
		{
			name: "List models",
			operation: func(ctx context.Context) (interface{}, error) {
				return manager.List(ctx)
			},
			validate: func(t *testing.T, result interface{}) {
				resp := result.(*api.ModelListResponse)
				assert.NotNil(t, resp.Models, "Models list should not be nil")
				assert.NotEmpty(t, resp.Models, "Should have at least one model")
				for _, model := range resp.Models {
					assert.NotEmpty(t, model.Name, "Model name should not be empty")
					assert.NotZero(t, model.Size, "Model size should not be zero")
				}
			},
			expectError: false,
		},
		{
			name: "Show model",
			operation: func(ctx context.Context) (interface{}, error) {
				return manager.Show(ctx, "llama3")
			},
			validate: func(t *testing.T, result interface{}) {
				resp := result.(*api.ShowModelResponse)
				assert.NotEmpty(t, resp.Name, "Model name should not be empty")
				assert.NotEmpty(t, resp.Description, "Model description should not be empty")
			},
			expectError: false,
		},
		{
			name: "Pull model",
			operation: func(ctx context.Context) (interface{}, error) {
				return nil, manager.Pull(ctx, "llama3:latest")
			},
			validate: func(t *testing.T, result interface{}) {
				// No result to validate for pull operation
			},
			expectError: false,
		},
		{
			name: "Push model",
			operation: func(ctx context.Context) (interface{}, error) {
				return nil, manager.Push(ctx, "llama3:latest")
			},
			validate: func(t *testing.T, result interface{}) {
				// No result to validate for push operation
			},
			expectError: false,
		},
		{
			name: "Copy model",
			operation: func(ctx context.Context) (interface{}, error) {
				return nil, manager.Copy(ctx, "llama3:latest", "llama3:copy")
			},
			validate: func(t *testing.T, result interface{}) {
				// No result to validate for copy operation
			},
			expectError: false,
		},
		{
			name: "Delete model",
			operation: func(ctx context.Context) (interface{}, error) {
				return nil, manager.Delete(ctx, "llama3:copy")
			},
			validate: func(t *testing.T, result interface{}) {
				// No result to validate for delete operation
			},
			expectError: false,
		},
		{
			name: "Create model",
			operation: func(ctx context.Context) (interface{}, error) {
				return nil, manager.Create(ctx, "llama3:test", "test-owner")
			},
			validate: func(t *testing.T, result interface{}) {
				// No result to validate for create operation
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the operation
			result, err := tt.operation(ctx)

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.validate != nil && result != nil {
				tt.validate(t, result)
			}
		})
	}
}
