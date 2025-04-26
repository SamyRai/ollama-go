package ollama

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCreation(t *testing.T) {
	// Test creating a new client with default configuration
	client := New()
	require.NotNil(t, client, "Client should not be nil")

	// Test the client has a raw client
	rawClient := client.RawClient()
	require.NotNil(t, rawClient, "Raw client should not be nil")

	// Test the client has a default HTTP client
	require.NotNil(t, rawClient.HTTPClient, "HTTP client should not be nil")
}

func TestClientConfiguration(t *testing.T) {
	// Test cases for client configuration
	tests := []struct {
		name      string
		configure func(*Client) *Client
		validate  func(*testing.T, *Client)
	}{
		{
			name: "WithBaseURL",
			configure: func(c *Client) *Client {
				return c.WithBaseURL("http://localhost:12345")
			},
			validate: func(t *testing.T, c *Client) {
				assert.Equal(t, "http://localhost:12345", c.config.BaseURL, "Base URL should be set")
			},
		},
		{
			name: "WithTimeout",
			configure: func(c *Client) *Client {
				return c.WithTimeout(30 * time.Second)
			},
			validate: func(t *testing.T, c *Client) {
				assert.Equal(t, 30*time.Second, c.config.Timeout, "Timeout should be set")
			},
		},
		{
			name: "WithAPIKey",
			configure: func(c *Client) *Client {
				return c.WithAPIKey("test-api-key")
			},
			validate: func(t *testing.T, c *Client) {
				assert.Equal(t, "test-api-key", c.config.APIKey, "API key should be set")
			},
		},
		{
			name: "WithHTTPClient",
			configure: func(c *Client) *Client {
				httpClient := &http.Client{
					Timeout: 45 * time.Second,
				}
				return c.WithHTTPClient(httpClient)
			},
			validate: func(t *testing.T, c *Client) {
				assert.Equal(t, 45*time.Second, c.client.HTTPClient.Timeout, "HTTP client timeout should be set")
			},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new client
			client := New()

			// Configure the client
			client = tt.configure(client)

			// Validate the configuration
			tt.validate(t, client)
		})
	}
}

func TestClientBuilders(t *testing.T) {
	// Create a new client
	client := New()

	// Test that all builder methods return non-nil builders
	assert.NotNil(t, client.Chat(), "Chat builder should not be nil")
	assert.NotNil(t, client.Completion(), "Completion builder should not be nil")
	assert.NotNil(t, client.Embeddings(), "Embeddings builder should not be nil")
	assert.NotNil(t, client.Models(), "Model manager should not be nil")
	assert.NotNil(t, client.Status(), "Status manager should not be nil")
}

func TestVersion(t *testing.T) {
	// Test that the version function returns a non-empty string
	version := Version()
	assert.NotEmpty(t, version, "Version should not be empty")
	assert.Contains(t, version, "ollama-go", "Version should contain 'ollama-go'")
}

func TestToolRegistry(t *testing.T) {
	// Test creating a new tool registry
	registry := NewToolRegistry()
	assert.NotNil(t, registry, "Tool registry should not be nil")
}

func TestOptions(t *testing.T) {
	// Test creating new options
	options := NewOptions()
	assert.NotNil(t, options, "Options should not be nil")

	// Test applying options
	ApplyOptions(options,
		WithTemperature(0.7),
		WithTopP(0.9),
	)

	assert.Equal(t, 0.7, options.Temperature, "Temperature should be set")
	assert.Equal(t, 0.9, options.TopP, "TopP should be set")
}
