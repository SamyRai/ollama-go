package tests

import (
	"context"
	"testing"

	"github.com/SamyRai/ollama-go/api"
	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/SamyRai/ollama-go/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatBuilder(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name        string
		builder     func() *api.ChatBuilder
		expectError bool
	}{
		{
			name: "Basic chat",
			builder: func() *api.ChatBuilder {
				return api.NewChatBuilder(cli).
					WithModel("llama3").
					WithMessage("user", "What is AI?")
			},
			expectError: false,
		},
		{
			name: "Chat with system message",
			builder: func() *api.ChatBuilder {
				return api.NewChatBuilder(cli).
					WithModel("llama3").
					WithSystemMessage("You are a helpful assistant.").
					WithMessage("user", "What is AI?")
			},
			expectError: false,
		},
		{
			name: "Chat with temperature",
			builder: func() *api.ChatBuilder {
				return api.NewChatBuilder(cli).
					WithModel("llama3").
					WithMessage("user", "What is AI?").
					WithTemperature(0.7)
			},
			expectError: false,
		},
		{
			name: "Chat with multiple messages",
			builder: func() *api.ChatBuilder {
				return api.NewChatBuilder(cli).
					WithModel("llama3").
					WithMessage("user", "Hello").
					WithMessage("assistant", "Hi there! How can I help you?").
					WithMessage("user", "What is AI?")
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the chat request
			builder := tt.builder()

			// Execute the request
			resp, err := builder.Execute(ctx)

			// Validate the result
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.NotEmpty(t, resp.Message.Content)
		})
	}
}

func TestChatBuilderWithTools(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Create a weather tool
	weatherTool := api.Tool{
		Type: "function",
		Function: api.ToolFunction{
			Name:        "getWeather",
			Description: "Get the current weather for a location",
			Parameters: map[string]api.ToolParam{
				"location": {
					Type:        "string",
					Description: "The city and state, e.g. San Francisco, CA",
				},
			},
		},
	}

	// Build the chat request
	builder := api.NewChatBuilder(cli).
		WithModel("llama3").
		WithMessage("user", "What's the weather like in Paris?").
		WithTools(weatherTool)

	// Execute the request
	resp, err := builder.Execute(ctx)

	// Validate the result
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Message.Content)

	// Check for tool calls
	if len(resp.Message.ToolCalls) > 0 {
		toolCall := resp.Message.ToolCalls[0]
		assert.Equal(t, "getWeather", toolCall.Function.Name)
		location, ok := toolCall.Function.Arguments["location"]
		assert.True(t, ok, "Tool call should have location argument")
		assert.Equal(t, "Paris", location)
	}
}

func TestChatBuilderStreaming(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Build the chat request
	builder := api.NewChatBuilder(cli).
		WithModel("llama3").
		WithMessage("user", "Write a short poem about Go programming.")

	// Track streaming responses
	var responses []*api.ChatResponse

	// Execute the streaming request
	err := builder.Stream(ctx, func(resp *api.ChatResponse) {
		responses = append(responses, resp)
	})

	// Validate the result
	require.NoError(t, err)
	assert.NotEmpty(t, responses, "Should receive streaming responses")

	// Validate the content
	var fullContent string
	for _, resp := range responses {
		fullContent += resp.Message.Content
	}
	assert.NotEmpty(t, fullContent, "Combined content should not be empty")
}
