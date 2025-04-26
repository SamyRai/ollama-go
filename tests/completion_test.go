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

// TestGenerateCompletion validates text completion API behavior.
func TestGenerateCompletion(t *testing.T) {
	rec, err := recorder.New("fixtures/generate_completion")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	req := structures.CompletionRequest{
		Model:  "llama3.1",
		Prompt: "Hello, world!",
		Stream: false,
	}

	resp, err := cli.GenerateCompletion(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "llama3.1", resp.Model)
	assert.NotEmpty(t, resp.Response)
}

// TestGenerateCompletionStream validates streaming text completion.
func TestGenerateCompletionStream(t *testing.T) {
	rec, err := recorder.New("fixtures/generate_completion_stream")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	req := structures.CompletionRequest{
		Model:  "llama3.1",
		Prompt: "Stream test",
		Stream: true,
	}

	resp, err := cli.GenerateCompletion(req)
	log.Printf("resp.Response: %v", resp)
	require.NoError(t, err)
	assert.NotEmpty(t, resp)
}
