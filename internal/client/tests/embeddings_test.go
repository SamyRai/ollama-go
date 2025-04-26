package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbeddings(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name        string
		fixture     string
		input       string
		expectError bool
	}{
		{
			name:    "Basic embedding request",
			fixture: "embeddings",
			input:   "The quick brown fox jumps over the lazy dog",
		},
		{
			name:    "Embedding request with longer text",
			fixture: "embeddings_long",
			input:   "Artificial intelligence (AI) is intelligence demonstrated by machines, as opposed to intelligence displayed by animals and humans. Example tasks in which this is done include speech recognition, computer vision, translation between (natural) languages, as well as other mappings of inputs.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup VCR
			client, cleanup := testutils.SetupVCR(t, tt.fixture)
			defer cleanup()

			// Create the request
			req := testutils.CreateEmbeddingRequest(tt.input)

			// Execute the request
			resp, err := client.GenerateEmbeddings(req)

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			testutils.AssertEmbeddingResponse(t, resp)

			// Additional assertions specific to embeddings
			assert.Equal(t, 1, len(resp.Embeddings), "Should have one embedding vector")
			assert.Greater(t, len(resp.Embeddings[0]), 0, "Embedding vector should have dimensions")
		})
	}
}

func TestEmbeddingsMultipleInputs(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "embeddings_multiple")
	defer cleanup()

	// Create the request with multiple inputs
	req := testutils.CreateEmbeddingRequest("")
	req.Input = []string{
		"The quick brown fox jumps over the lazy dog",
		"The five boxing wizards jump quickly",
	}

	// Execute the request
	resp, err := client.GenerateEmbeddings(req)

	// Validate the result
	require.NoError(t, err)
	testutils.AssertEmbeddingResponse(t, resp)

	// Additional assertions for multiple inputs
	assert.Equal(t, 2, len(resp.Embeddings), "Should have two embedding vectors")
	assert.Equal(t, len(resp.Embeddings[0]), len(resp.Embeddings[1]), "Embedding vectors should have the same dimensions")
}
