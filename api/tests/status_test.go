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

func TestStatusManager(t *testing.T) {
	// Create a context
	ctx := context.Background()

	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create the status manager
	manager := api.NewStatusManager(cli)

	// Test cases for status operations
	tests := []struct {
		name        string
		operation   func(ctx context.Context) (interface{}, error)
		validate    func(t *testing.T, result interface{})
		expectError bool
	}{
		{
			name: "Get version",
			operation: func(ctx context.Context) (interface{}, error) {
				return manager.GetVersion(ctx)
			},
			validate: func(t *testing.T, result interface{}) {
				resp := result.(*api.VersionResponse)
				assert.NotEmpty(t, resp.Version, "Version should not be empty")
			},
			expectError: false,
		},
		{
			name: "Get running processes",
			operation: func(ctx context.Context) (interface{}, error) {
				return manager.GetRunningProcesses(ctx)
			},
			validate: func(t *testing.T, result interface{}) {
				resp := result.(*api.ModelProcessResponse)
				assert.NotNil(t, resp.Models, "Models field should be initialized")

				// If there are running models, validate their properties
				for _, model := range resp.Models {
					assert.NotEmpty(t, model.Name, "Model name should not be empty")
					assert.NotEmpty(t, model.Model, "Model identifier should not be empty")
				}
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
