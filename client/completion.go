package client

import (
	"encoding/json"
	"github.com/SamyRai/ollama-go/structures"
)

// GenerateCompletion handles both streaming and non-streaming text generation.
func (c *OllamaClient) GenerateCompletion(req structures.CompletionRequest, callback func(structures.CompletionResponse)) (*structures.CompletionResponse, error) {
	if req.Stream {
		// Handle streaming response
		return nil, c.StreamRequest("POST", "/api/generate", req, func(data json.RawMessage) {
			var completionResp structures.CompletionResponse
			if err := json.Unmarshal(data, &completionResp); err == nil {
				callback(completionResp)
			}
		})
	}

	// Handle normal response
	var resp structures.CompletionResponse
	err := c.Request("POST", "/api/generate", req, &resp)
	return &resp, err
}
