package api

import (
	"context"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/structures"
)

// EmbeddingResponse contains generated embeddings.
type EmbeddingResponse = structures.EmbeddingResponse

// EmbeddingsBuilder provides a fluent interface for building embeddings requests.
type EmbeddingsBuilder struct {
	client   *client.OllamaClient
	model    string
	input    []string
	truncate bool
	options  structures.Options
}

// NewEmbeddingsBuilder creates a new EmbeddingsBuilder.
func NewEmbeddingsBuilder(client *client.OllamaClient) *EmbeddingsBuilder {
	return &EmbeddingsBuilder{
		client: client,
		input:  make([]string, 0),
	}
}

// WithModel sets the model to use for the embeddings.
func (b *EmbeddingsBuilder) WithModel(model string) *EmbeddingsBuilder {
	b.model = model
	return b
}

// WithInput adds input text to the embeddings request.
func (b *EmbeddingsBuilder) WithInput(input string) *EmbeddingsBuilder {
	b.input = append(b.input, input)
	return b
}

// WithInputs sets all input texts for the embeddings request.
func (b *EmbeddingsBuilder) WithInputs(inputs []string) *EmbeddingsBuilder {
	b.input = inputs
	return b
}

// WithTruncate sets whether to truncate input if needed.
func (b *EmbeddingsBuilder) WithTruncate(truncate bool) *EmbeddingsBuilder {
	b.truncate = truncate
	return b
}

// WithOptions sets all options for the embeddings.
func (b *EmbeddingsBuilder) WithOptions(options structures.Options) *EmbeddingsBuilder {
	b.options = options
	return b
}

// Execute sends the embeddings request and returns the response.
func (b *EmbeddingsBuilder) Execute(_ context.Context) (*EmbeddingResponse, error) {
	req := structures.EmbeddingRequest{
		Model:    b.model,
		Input:    b.input,
		Truncate: b.truncate,
		Options:  b.options,
	}

	return b.client.GenerateEmbeddings(req)
}
