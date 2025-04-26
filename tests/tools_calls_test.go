package tests

import (
	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/SamyRai/ollama-go/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
	"log"
	"testing"
)

// TestChat validates the chat API functionality and tool call management.
func TestChatWithTools(t *testing.T) {
	// Start the VCR recorder to mock the HTTP requests
	rec, err := recorder.New("fixtures/chat_with_tools")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	// Define the tools available to the model during chat
	tools := []structures.Tool{
		{
			Type: "function",
			Function: structures.ToolFunction{
				Name:        "getWeather",
				Description: "Retrieve the weather",
				Parameters: map[string]structures.ToolParam{
					"location": {
						Type:        "string",
						Description: "The location for the weather forecast",
					},
				},
			},
		},
	}

	// Construct the chat request with the tools
	req := structures.ChatRequest{
		Model: "llama3.1",
		Messages: []structures.Message{
			{Role: "user", Content: "What is the weather like in Paris?"},
		},
		Tools: tools,
	}

	// Make the API call to Chat
	resp, err := cli.Chat(req)

	// Validate no errors occurred during the request
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Assert that the response message is as expected
	assert.Equal(t, "assistant", resp.Message.Role)
	log.Print(resp)
	log.Printf("resp.Message.Content: %s", resp.Message.Content)
	log.Printf("resp.Message.Tools_calls: %s", resp.Message.ToolCalls)
	log.Printf("resp.Tools_calls: %s", resp.ToolCalls)

	// Check that tool calls were correctly passed and returned in the response
	require.NotNil(t, resp.Message.ToolCalls)
	assert.Len(t, resp.Message.ToolCalls, 1)

	// Ensure the tool call details are correct
	assert.Equal(t, "getWeather", resp.Message.ToolCalls[0].Function.Name)
	assert.Equal(t, "Paris", resp.Message.ToolCalls[0].Function.Arguments["location"])
}
