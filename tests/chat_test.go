package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
	"hrelay/core/llm/ollama/client"
	"hrelay/core/llm/ollama/config"
	"hrelay/core/llm/ollama/structures"
	"testing"
)

// TestChat validates the chat API functionality.
func TestChat(t *testing.T) {
	rec, err := recorder.New("fixtures/chat")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	req := structures.ChatRequest{
		Model: "llama3.1",
		Messages: []structures.Message{
			{Role: "user", Content: "What is AI?"},
		},
	}

	resp, err := cli.Chat(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "assistant", resp.Message.Role)
	assert.NotEmpty(t, resp.Message.Content)
}
