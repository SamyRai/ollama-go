package client

import (
	"encoding/json"

	"github.com/SamyRai/ollama-go/internal/structures"
)

// Chat handles both streaming and non-streaming chat interactions.
func (c *OllamaClient) Chat(req structures.ChatRequest, callback func(structures.ChatResponse)) (*structures.ChatResponse, error) {
	if req.Stream {
		// Handle streaming response
		return nil, c.StreamRequest("POST", "/api/chat", req, func(data json.RawMessage) {
			var chatResp structures.ChatResponse
			if err := json.Unmarshal(data, &chatResp); err == nil {
				callback(chatResp)
			}
		})
	}

	// Handle normal response
	var resp structures.ChatResponse
	err := c.Request("POST", "/api/chat", req, &resp)
	return &resp, err
}
