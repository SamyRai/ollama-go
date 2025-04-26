package client

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// GenerateEmbeddings retrieves text embeddings from the API.
func (c *OllamaClient) GenerateEmbeddings(req structures.EmbeddingRequest) (*structures.EmbeddingResponse, error) {
	var resp structures.EmbeddingResponse
	err := c.Request("POST", "/api/embed", req, &resp)
	return &resp, err
}
