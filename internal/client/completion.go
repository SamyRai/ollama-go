package client

import (
	"encoding/json"

	"github.com/SamyRai/ollama-go/internal/structures"
)

// GenerateCompletion handles both streaming and non-streaming text generation.
func (c *OllamaClient) GenerateCompletion(req structures.CompletionRequest, callback func(structures.CompletionResponse)) (*structures.CompletionResponse, error) {
	c.Logger.Debug("Processing completion request to model: %s", req.Model)

	if req.Stream {
		// Handle streaming response
		c.Logger.Debug("Using streaming mode for completion request")
		return nil, c.StreamRequest("POST", "/api/generate", req, func(data json.RawMessage) {
			var completionResp structures.CompletionResponse
			if err := json.Unmarshal(data, &completionResp); err == nil {
				c.Logger.Debug("Received completion response chunk")
				callback(completionResp)
			} else {
				c.Logger.Error("Failed to unmarshal completion response: %v", err)
			}
		})
	}

	// Handle normal response
	c.Logger.Debug("Using non-streaming mode for completion request")
	var resp structures.CompletionResponse
	err := c.Request("POST", "/api/generate", req, &resp)
	if err != nil {
		c.Logger.Error("Completion request failed: %v", err)
		return nil, err
	}

	c.Logger.Debug("Successfully completed text generation request")
	return &resp, err
}
