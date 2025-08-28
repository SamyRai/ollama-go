package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	assert.ErrorContains(t, err, "no such host")
}

// TestInvalidEndpoint ensures the client handles 404 errors correctly.
func TestInvalidEndpoint(t *testing.T) {
	// Create a mock server that always returns 404
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	// Create a client that points to the mock server
	cli := client.NewClient(&config.Config{
		BaseURL: server.URL,
	})

	var resp interface{}
	err := cli.Request("GET", "/api/invalid-endpoint", nil, &resp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "404", "Expected a 404 error for an invalid endpoint")
}
