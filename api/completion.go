package api

import (
	"context"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/structures"
)

// CompletionResponse represents the model's response to a text generation request.
type CompletionResponse = structures.CompletionResponse

// CompletionBuilder provides a fluent interface for building completion requests.
type CompletionBuilder struct {
	client  *client.OllamaClient
	model   string
	prompt  string
	suffix  string
	images  []string
	options structures.Options
	stream  bool
	raw     bool
}

// NewCompletionBuilder creates a new CompletionBuilder.
func NewCompletionBuilder(client *client.OllamaClient) *CompletionBuilder {
	return &CompletionBuilder{
		client: client,
		images: make([]string, 0),
	}
}

// WithModel sets the model to use for the completion.
func (b *CompletionBuilder) WithModel(model string) *CompletionBuilder {
	b.model = model
	return b
}

// WithPrompt sets the prompt for the completion.
func (b *CompletionBuilder) WithPrompt(prompt string) *CompletionBuilder {
	b.prompt = prompt
	return b
}

// WithSuffix sets the suffix for the completion.
func (b *CompletionBuilder) WithSuffix(suffix string) *CompletionBuilder {
	b.suffix = suffix
	return b
}

// WithImages adds images to the completion request.
func (b *CompletionBuilder) WithImages(images ...string) *CompletionBuilder {
	b.images = append(b.images, images...)
	return b
}

// WithTemperature sets the temperature for the completion.
func (b *CompletionBuilder) WithTemperature(temperature float64) *CompletionBuilder {
	b.options.Temperature = temperature
	return b
}

// WithTopP sets the top_p value for the completion.
func (b *CompletionBuilder) WithTopP(topP float64) *CompletionBuilder {
	b.options.TopP = topP
	return b
}

// WithTopK sets the top_k value for the completion.
func (b *CompletionBuilder) WithTopK(topK int) *CompletionBuilder {
	b.options.TopK = topK
	return b
}

// WithOptions sets all options for the completion.
func (b *CompletionBuilder) WithOptions(options structures.Options) *CompletionBuilder {
	b.options = options
	return b
}

// WithRaw sets whether to return raw model output.
func (b *CompletionBuilder) WithRaw(raw bool) *CompletionBuilder {
	b.raw = raw
	return b
}

// Execute sends the completion request and returns the response.
func (b *CompletionBuilder) Execute(ctx context.Context) (*CompletionResponse, error) {
	req := structures.CompletionRequest{
		Model:   b.model,
		Prompt:  b.prompt,
		Suffix:  b.suffix,
		Images:  b.images,
		Options: b.options,
		Stream:  false,
		Raw:     b.raw,
	}

	// Use a no-op callback since we're not streaming
	resp, err := b.client.GenerateCompletion(req, func(resp structures.CompletionResponse) {})
	return resp, err
}

// Stream sends the completion request and streams the response through the callback.
func (b *CompletionBuilder) Stream(ctx context.Context, callback func(*CompletionResponse)) error {
	req := structures.CompletionRequest{
		Model:   b.model,
		Prompt:  b.prompt,
		Suffix:  b.suffix,
		Images:  b.images,
		Options: b.options,
		Stream:  true,
		Raw:     b.raw,
	}

	_, err := b.client.GenerateCompletion(req, func(resp structures.CompletionResponse) {
		callback(&resp)
	})
	return err
}
