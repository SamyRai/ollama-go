package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateCompletion validates text completion API behavior.
func TestGenerateCompletion(t *testing.T) {
	cli, rec := SetupVCRTest(t, "generate_completion")
	defer func() {
		err := rec.Stop()
		if err != nil {
			t.Logf("Failed to stop recorder: %v", err)
		}
	}()

	req := structures.CompletionRequest{
		Model:  "llama3.1", // Match the model name used in the recorded fixture
		Prompt: "Hello, world!",
		Stream: false,
	}

	resp, err := cli.GenerateCompletion(req, func(_ structures.CompletionResponse) {
		// Streaming callback - not needed for non-streaming test
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "llama3.1", resp.Model) // Match the model name in the response
	assert.NotEmpty(t, resp.Response)
}

// TestGenerateCompletionStream validates streaming text completion.
func TestGenerateCompletionStream(t *testing.T) {
	cli, rec := SetupVCRTest(t, "generate_completion_stream")
	defer func() {
		err := rec.Stop()
		if err != nil {
			t.Logf("Failed to stop recorder: %v", err)
		}
	}()

	req := structures.CompletionRequest{
		Model:  "llama3.1", // Match the model name used in the recorded fixture
		Prompt: "Stream test",
		Stream: true,
	}

	// For streaming tests, we need to collect all responses
	var responses []structures.CompletionResponse

	// Execute the streaming request
	resp, err := cli.GenerateCompletion(req, func(response structures.CompletionResponse) {
		responses = append(responses, response)
		t.Log("Received streaming response chunk")
	})

	require.NoError(t, err)

	// For streaming requests, we should validate we got streaming chunks
	// Note: The final response might be nil for streaming requests depending on implementation
	if resp == nil {
		t.Log("Final response is nil, which is expected for some streaming implementations")
		require.NotEmpty(t, responses, "Should have received at least some streaming responses")
	} else {
		require.NotNil(t, resp)
		assert.NotEmpty(t, resp.Response)
	}
}
