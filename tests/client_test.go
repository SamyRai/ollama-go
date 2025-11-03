package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestOllamaClientInitialization checks if the client initializes correctly.
func TestOllamaClientInitialization(t *testing.T) {
	cli, _ := SetupVCRTest(t, "client_test")
	require.NotNil(t, cli)
	assert.Equal(t, "http://localhost:11434", cli.BaseURL)
	assert.NotNil(t, cli.HTTPClient)
}

// TestRequestFailure ensures the client handles failed requests properly.
func TestRequestFailure(t *testing.T) {
	cli := client.NewClient(&config.Config{
		BaseURL: "http://localhostbroke:11434",
	})

	var resp interface{}
	err := cli.Request("GET", "/api/version", nil, &resp)
	require.Error(t, err, "Expected an error for an invalid URL")
	assert.Nil(t, resp, "Response should be nil when request fails")
}

// TestInvalidEndpoint ensures the client handles 404 errors correctly.
func TestInvalidEndpoint(t *testing.T) {
	t.Skip("Skipping test that requires a running Ollama server to re-record cassettes.")
	cli, _ := SetupVCRTest(t, "client_test")

	var resp interface{}
	err := cli.Request("GET", "/api/invalid-endpoint", nil, &resp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "404", "Expected a 404 error for an invalid endpoint")
}
