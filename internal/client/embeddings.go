package client

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// GenerateEmbeddings retrieves text embeddings from the API.
func (c *OllamaClient) GenerateEmbeddings(req structures.EmbeddingRequest) (*structures.EmbeddingResponse, error) {
	c.Logger.Debug("Processing embeddings request for model: %s", req.Model)

	var resp structures.EmbeddingResponse
	err := c.Request("POST", "/api/embed", req, &resp)
	if err != nil {
		c.Logger.Error("Failed to generate embeddings: %v", err)
		return nil, err
	}

	embeddingCount := 0
	if len(resp.Embeddings) > 0 {
		embeddingCount = len(resp.Embeddings)
	}
	c.Logger.Debug("Successfully generated embeddings, count: %d", embeddingCount)
	return &resp, nil
}
