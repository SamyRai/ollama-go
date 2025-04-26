package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVersion(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "get_version")
	defer cleanup()

	// Execute the request
	resp, err := client.GetVersion()

	// Validate the result
	require.NoError(t, err)
	assert.NotNil(t, resp, "Response should not be nil")
	assert.NotEmpty(t, resp.Version, "Version should not be empty")
}

func TestGetRunningProcesses(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "get_running_processes")
	defer cleanup()

	// Execute the request
	resp, err := client.GetRunningProcesses()

	// Validate the result
	require.NoError(t, err)
	assert.NotNil(t, resp, "Response should not be nil")

	// Note: There might not be any running models, so we don't assert on the length
	// But we can check that the models field is initialized
	assert.NotNil(t, resp.Models, "Models field should be initialized")

	// If there are running models, validate their properties
	for _, model := range resp.Models {
		assert.NotEmpty(t, model.Name, "Model name should not be empty")
		assert.NotEmpty(t, model.Model, "Model identifier should not be empty")
		assert.NotZero(t, model.Size, "Model size should not be zero")
	}
}
