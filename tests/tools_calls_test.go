package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestChat validates the chat API functionality and tool call management.
func TestChatWithTools(t *testing.T) {
	cli, rec := SetupVCRTest(t, "chat_with_tools")
	defer rec.Stop()

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
		Model: "llama3", // Make sure model name matches what was used in recording
		Messages: []structures.Message{
			{Role: "user", Content: "What is the weather like in Paris?"},
		},
		Tools: tools,
	}

	// Make the API call to Chat
	resp, err := cli.Chat(req, func(response structures.ChatResponse) {
		// Callback for streaming responses - not needed for this test
	})

	// Validate no errors occurred during the request
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Assert that the response message is as expected
	assert.Equal(t, "assistant", resp.Message.Role)

	// Check that tool calls were correctly passed and returned in the response
	require.NotNil(t, resp.Message.ToolCalls)
	assert.Len(t, resp.Message.ToolCalls, 1)

	// Ensure the tool call details are correct
	assert.Equal(t, "getWeather", resp.Message.ToolCalls[0].Function.Name)
	assert.Equal(t, "Paris", resp.Message.ToolCalls[0].Function.Arguments["location"])
}
