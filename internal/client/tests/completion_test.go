package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompletion(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name        string
		fixture     string
		request     structures.CompletionRequest
		expectError bool
	}{
		{
			name:    "Basic completion request",
			fixture: "generate_completion",
			request: testutils.CreateCompletionRequest("Once upon a time"),
		},
		{
			name:    "Completion request with options",
			fixture: "generate_completion_with_options",
			request: func() structures.CompletionRequest {
				req := testutils.CreateCompletionRequest("List 5 benefits of AI:")
				req.Options.Temperature = 0.7
				req.Options.TopP = 0.9
				return req
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup VCR
			client, cleanup := testutils.SetupVCR(t, tt.fixture)
			defer cleanup()

			// Execute the request
			resp, err := client.GenerateCompletion(tt.request, func(structures.CompletionResponse) {})

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			testutils.AssertCompletionResponse(t, resp)
		})
	}
}

func TestCompletionStreaming(t *testing.T) {
	// Setup VCR
	client, cleanup := testutils.SetupVCR(t, "generate_completion_stream")
	defer cleanup()

	// Create the request
	req := testutils.CreateCompletionRequest("Write a short story about a robot.")
	req.Stream = true

	// Track streaming responses
	var responses []structures.CompletionResponse

	// Execute the streaming request
	_, err := client.GenerateCompletion(req, func(resp structures.CompletionResponse) {
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
