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

func TestCompletionBuilder(t *testing.T) {
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
		builder     func() *api.CompletionBuilder
		expectError bool
	}{
		{
			name: "Basic completion",
			builder: func() *api.CompletionBuilder {
				return api.NewCompletionBuilder(cli).
					WithModel("llama3").
					WithPrompt("Once upon a time")
			},
			expectError: false,
		},
		{
			name: "Completion with temperature",
			builder: func() *api.CompletionBuilder {
				return api.NewCompletionBuilder(cli).
					WithModel("llama3").
					WithPrompt("Once upon a time").
					WithTemperature(0.7)
			},
			expectError: false,
		},
		{
			name: "Completion with suffix",
			builder: func() *api.CompletionBuilder {
				return api.NewCompletionBuilder(cli).
					WithModel("llama3").
					WithPrompt("Once upon a time").
					WithSuffix(" in a land far away.")
			},
			expectError: false,
		},
		{
			name: "Completion with options",
			builder: func() *api.CompletionBuilder {
				options := api.NewOptions()
				api.ApplyOptions(options,
					api.WithTemperature(0.7),
					api.WithTopP(0.9),
					api.WithTopK(40),
				)
				return api.NewCompletionBuilder(cli).
					WithModel("llama3").
					WithPrompt("List 5 benefits of AI:").
					WithOptions(*options)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the completion request
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
			assert.NotEmpty(t, resp.Response)
		})
	}
}

func TestCompletionBuilderStreaming(t *testing.T) {
	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Build the completion request
	builder := api.NewCompletionBuilder(cli).
		WithModel("llama3").
		WithPrompt("Write a short story about a robot learning to paint.")

	// Track streaming responses
	var responses []*api.CompletionResponse

	// Execute the streaming request
	err := builder.Stream(ctx, func(resp *api.CompletionResponse) {
		responses = append(responses, resp)
	})

	// Validate the result
	require.NoError(t, err)
	assert.NotEmpty(t, responses, "Should receive streaming responses")

	// Validate the content
	var fullContent string
	for _, resp := range responses {
		fullContent += resp.Response
	}
	assert.NotEmpty(t, fullContent, "Combined content should not be empty")
}

func TestCompletionBuilderWithImages(t *testing.T) {
	// Skip this test if not running against a real server
	t.Skip("This test requires a real server with multimodal capabilities")

	// Setup mock server
	mockServer := testutils.NewMockServer()
	defer mockServer.Close()

	// Create a client with the mock server URL
	cfg := config.DefaultConfig()
	cfg.BaseURL = mockServer.URL()
	cli := client.NewClient(cfg)

	// Create a context
	ctx := context.Background()

	// Build the completion request with an image
	builder := api.NewCompletionBuilder(cli).
		WithModel("llava").
		WithPrompt("Describe this image:").
		WithImages("data:image/jpeg;base64,...")

	// Execute the request
	resp, err := builder.Execute(ctx)

	// Validate the result
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Response)
}
