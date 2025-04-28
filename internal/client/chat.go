package client

import (
	"encoding/json"

	"github.com/SamyRai/ollama-go/internal/structures"
)

// Chat handles both streaming and non-streaming chat interactions.
func (c *OllamaClient) Chat(req structures.ChatRequest, callback func(structures.ChatResponse)) (*structures.ChatResponse, error) {
	c.Logger.Debug("Processing chat request to model: %s", req.Model)

	if req.Stream {
		// Handle streaming response
		c.Logger.Debug("Using streaming mode for chat request")
		return nil, c.StreamRequest("POST", "/api/chat", req, func(data json.RawMessage) {
			var chatResp structures.ChatResponse
			if err := json.Unmarshal(data, &chatResp); err == nil {
				c.Logger.Debug("Received chat response chunk")
				callback(chatResp)
			} else {
				c.Logger.Error("Failed to unmarshal chat response: %v", err)
			}
		})
	}

	// Handle normal response
	c.Logger.Debug("Using non-streaming mode for chat request")
	var resp structures.ChatResponse
	err := c.Request("POST", "/api/chat", req, &resp)
	if err != nil {
		c.Logger.Error("Chat request failed: %v", err)
		return nil, err
	}

	c.Logger.Debug("Successfully completed chat request")
	return &resp, err
}
