package testutils

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertChatResponse validates common properties of a chat response.
func AssertChatResponse(t *testing.T, resp *structures.ChatResponse) {
	t.Helper()

	require.NotNil(t, resp, "Response should not be nil")
	assert.NotEmpty(t, resp.Model, "Response model should not be empty")
	assert.NotZero(t, resp.CreatedAt, "Response creation time should not be zero")
	assert.NotEmpty(t, resp.Message.Role, "Response message role should not be empty")
	assert.NotEmpty(t, resp.Message.Content, "Response message content should not be empty")
}

// AssertCompletionResponse validates common properties of a completion response.
func AssertCompletionResponse(t *testing.T, resp *structures.CompletionResponse) {
	t.Helper()

	require.NotNil(t, resp, "Response should not be nil")
	assert.NotEmpty(t, resp.Model, "Response model should not be empty")
	assert.NotZero(t, resp.CreatedAt, "Response creation time should not be zero")
	assert.NotEmpty(t, resp.Response, "Response text should not be empty")
}

// AssertEmbeddingResponse validates common properties of an embedding response.
func AssertEmbeddingResponse(t *testing.T, resp *structures.EmbeddingResponse) {
	t.Helper()

	require.NotNil(t, resp, "Response should not be nil")
	assert.NotEmpty(t, resp.Model, "Response model should not be empty")
	assert.NotEmpty(t, resp.Embeddings, "Embeddings should not be empty")

	// Check that embeddings have reasonable dimensions
	if len(resp.Embeddings) > 0 {
		assert.True(t, len(resp.Embeddings[0]) > 0, "Embedding vector should have dimensions")
	}
}

// AssertModelListResponse validates common properties of a model list response.
func AssertModelListResponse(t *testing.T, resp *structures.ModelListResponse) {
	t.Helper()

	require.NotNil(t, resp, "Response should not be nil")
	assert.NotNil(t, resp.Models, "Models list should not be nil")
}

// AssertToolCall validates a tool call.
func AssertToolCall(t *testing.T, toolCall structures.ToolCall, expectedName string) {
	t.Helper()

	assert.Equal(t, expectedName, toolCall.Function.Name, "Tool call function name should match")
	assert.NotNil(t, toolCall.Function.Arguments, "Tool call arguments should not be nil")
}
