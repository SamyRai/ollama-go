package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListModels(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "list_models")
	defer cleanup()

	// Execute the request
	resp, err := client.ListModels()

	// Validate the result
	require.NoError(t, err)
	testutils.AssertModelListResponse(t, resp)

	// Additional assertions
	assert.NotEmpty(t, resp.Models, "Should have at least one model")
	for _, model := range resp.Models {
		assert.NotEmpty(t, model.Name, "Model name should not be empty")
		assert.NotZero(t, model.Size, "Model size should not be zero")
	}
}

func TestShowModel(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "show_model")
	defer cleanup()

	// Create the request
	req := structures.ShowModelRequest{
		Model: testutils.TestModel,
	}

	// Execute the request
	resp, err := client.ShowModel(req)

	// Validate the result
	require.NoError(t, err)
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testutils.TestModel, resp.Name, "Model name should match")
	assert.NotEmpty(t, resp.Description, "Model description should not be empty")
}

func TestModelManagement(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name        string
		fixture     string
		operation   string
		modelName   string
		expectError bool
	}{
		{
			name:      "Pull model",
			fixture:   "pull_model",
			operation: "pull",
			modelName: "llama3:latest",
		},
		{
			name:      "Push model",
			fixture:   "push_model",
			operation: "push",
			modelName: "llama3:latest",
		},
		{
			name:      "Delete model",
			fixture:   "delete_model",
			operation: "delete",
			modelName: "llama3:test",
		},
		{
			name:      "Copy model",
			fixture:   "copy_model",
			operation: "copy",
			modelName: "llama3:latest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup VCR
			client, cleanup := testutils.SetupVCR(t, tt.fixture)
			defer cleanup()

			var err error

			// Execute the appropriate operation
			switch tt.operation {
			case "pull":
				err = client.PullModel(tt.modelName)
			case "push":
				err = client.PushModel(tt.modelName)
			case "delete":
				err = client.DeleteModel(tt.modelName)
			case "copy":
				err = client.CopyModel(tt.modelName, tt.modelName+"-copy")
			case "create":
				err = client.CreateModel(structures.ModelManagementRequest{
					Name: tt.modelName,
				})
			}

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
