// Package tests contains integration tests for the Ollama Go client API.
// Tests use the VCR library to record API interactions for repeatable testing.
package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestChat validates the chat API functionality.
func TestChat(t *testing.T) {
	cli, rec := SetupVCRTest(t, "chat")
	defer func() {
		err := rec.Stop()
		if err != nil {
			t.Logf("Failed to stop recorder: %v", err)
		}
	}()

	req := structures.ChatRequest{
		Model: "llama3", // Make sure model name matches the one used when recording
		Messages: []structures.Message{
			{Role: "user", Content: "What is AI?"},
		},
	}

	resp, err := cli.Chat(req, func(_ structures.ChatResponse) {
		// Streaming callback - not needed for this test
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "assistant", resp.Message.Role)
	assert.NotEmpty(t, resp.Message.Content)
}
