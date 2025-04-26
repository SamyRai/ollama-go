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

func TestEmbeddingsBuilder(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name        string
		builder     func() *api.EmbeddingsBuilder
		expectError bool
	}{
		{
			name: "Basic embeddings",
			builder: func() *api.EmbeddingsBuilder {
				return api.NewEmbeddingsBuilder(cli).
					WithModel("llama3").
					WithInput("The quick brown fox jumps over the lazy dog")
			},
			expectError: false,
		},
		{
			name: "Embeddings with truncate",
			builder: func() *api.EmbeddingsBuilder {
				return api.NewEmbeddingsBuilder(cli).
					WithModel("llama3").
					WithInput("The quick brown fox jumps over the lazy dog").
					WithTruncate(true)
			},
			expectError: false,
		},
		{
			name: "Embeddings with options",
			builder: func() *api.EmbeddingsBuilder {
				options := api.NewOptions()
				api.ApplyOptions(options,
					api.WithTemperature(0.0), // Temperature doesn't affect embeddings, but testing the API
				)
				return api.NewEmbeddingsBuilder(cli).
					WithModel("llama3").
					WithInput("The quick brown fox jumps over the lazy dog").
					WithOptions(*options)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the embeddings request
			builder := tt.builder()

			// Execute the request
			resp, err := builder.Execute(ctx)

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.NotEmpty(t, resp.Embeddings, "Embeddings should not be empty")
			assert.Equal(t, 1, len(resp.Embeddings), "Should have one embedding vector")
			assert.Greater(t, len(resp.Embeddings[0]), 0, "Embedding vector should have dimensions")
		})
	}
}

func TestEmbeddingsBuilderMultipleInputs(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Build the embeddings request with multiple inputs
	builder := api.NewEmbeddingsBuilder(cli).
		WithModel("llama3").
		WithInputs([]string{
			"The quick brown fox jumps over the lazy dog",
			"The five boxing wizards jump quickly",
		})

	// Execute the request
	resp, err := builder.Execute(ctx)

	// Validate the result
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Embeddings, "Embeddings should not be empty")
	assert.Equal(t, 2, len(resp.Embeddings), "Should have two embedding vectors")
	assert.Greater(t, len(resp.Embeddings[0]), 0, "First embedding vector should have dimensions")
	assert.Greater(t, len(resp.Embeddings[1]), 0, "Second embedding vector should have dimensions")
	assert.Equal(t, len(resp.Embeddings[0]), len(resp.Embeddings[1]), "Both embedding vectors should have the same dimensions")
}
