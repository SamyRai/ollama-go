package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChat(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name        string
		fixture     string
		request     structures.ChatRequest
		expectError bool
	}{
		{
			name:    "Basic chat request",
			fixture: "chat",
			request: testutils.CreateChatRequest("What is AI?"),
		},
		{
			name:    "Chat request with system message",
			fixture: "chat_with_system",
			request: testutils.CreateChatRequestWithSystemMessage(
				"You are a helpful assistant.",
				"What is AI?",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup VCR
			client, cleanup := testutils.SetupVCR(t, tt.fixture)
			defer cleanup()

			// Execute the request
			resp, err := client.Chat(tt.request, func(structures.ChatResponse) {})

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			testutils.AssertChatResponse(t, resp)
		})
	}
}

func TestChatStreaming(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "chat_stream")
	defer cleanup()

	// Create the request
	req := testutils.CreateChatRequest("Write a short poem about Go programming.")
	req.Stream = true

	// Track streaming responses
	var responses []structures.ChatResponse

	// Execute the streaming request
	_, err := client.Chat(req, func(resp structures.ChatResponse) {
		responses = append(responses, resp)
	})

	// Validate the result
	require.NoError(t, err)
	assert.NotEmpty(t, responses, "Should receive streaming responses")

	// Validate the last response
	if len(responses) > 0 {
		lastResp := responses[len(responses)-1]
		assert.True(t, lastResp.Done, "Last response should have Done=true")
	}
}
