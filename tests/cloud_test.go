package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go"
	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCloudClientInitialization checks if the client initializes correctly for cloud usage.
func TestCloudClientInitialization(t *testing.T) {
	cli := ollama.New().WithCloud().WithAPIKey("test-api-key")
	require.NotNil(t, cli)
}

// TestCloudChat validates the chat API functionality with the cloud service.
func TestCloudChat(t *testing.T) {
	t.Skip("Skipping test that requires a running Ollama server to re-record cassettes.")
	cli, _ := SetupVCRTest(t, "cloud_chat")

	// The request can be generic as the response is mocked.
	req := structures.ChatRequest{
		Model: "llama3",
		Messages: []structures.Message{
			{Role: "user", Content: "What is AI?"},
		},
	}

	resp, err := cli.Chat(req, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "assistant", resp.Message.Role)
	assert.NotEmpty(t, resp.Message.Content)
}
