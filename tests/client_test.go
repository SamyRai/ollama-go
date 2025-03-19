package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hrelay/core/llm/ollama/client"
	"hrelay/core/llm/ollama/config"
	"log"
	"testing"
)

// TestOllamaClientInitialization checks if the client initializes correctly.
func TestOllamaClientInitialization(t *testing.T) {
	cli := client.NewClient(config.DefaultConfig())
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
	log.Printf("Error: %v", err)
	log.Printf("Response: %v", resp)
	require.Error(t, err, "Expected an error for an invalid URL")
}

// TestInvalidEndpoint ensures the client handles 404 errors correctly.
func TestInvalidEndpoint(t *testing.T) {
	cli := client.NewClient(config.DefaultConfig())

	var resp interface{}
	err := cli.Request("GET", "/api/invalid-endpoint", nil, &resp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "404", "Expected a 404 error for an invalid endpoint")
}
