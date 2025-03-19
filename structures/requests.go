package structures

// =========================
// == Chat API ==
// =========================

// Message represents a single message in a chat conversation.
type Message struct {
	Role      string     `json:"role"`                 // Role of the sender (e.g., "assistant").
	Content   string     `json:"content"`              // Message content.
	Images    []string   `json:"images,omitempty"`     // Optional: Base64-encoded images (for multimodal models).
	ToolCalls []ToolCall `json:"tool_calls,omitempty"` // Optional: Tool calls made by the model.
}

// ChatRequest is used to engage in a chat conversation with a model.
type ChatRequest struct {
	Model     string    `json:"model"`                // Required: Name of the model to use.
	Messages  []Message `json:"messages"`             // Required: Chat history messages.
	Tools     []Tool    `json:"tools,omitempty"`      // Optional: Available tools.
	Format    string    `json:"format,omitempty"`     // Optional: Response format.
	Options   Options   `json:"options,omitempty"`    // Optional: Additional options.
	Stream    bool      `json:"stream,omitempty"`     // Optional: Whether to stream responses.
	KeepAlive string    `json:"keep_alive,omitempty"` // Optional: Duration to keep model in memory.
}

// =========================
// == Embeddings API ==
// =========================

// EmbeddingRequest is used to generate embeddings.
type EmbeddingRequest struct {
	Model     string   `json:"model"`             // Model name.
	Input     []string `json:"input"`             // Input text(s).
	Truncate  bool     `json:"truncate"`          // Whether to truncate input if needed.
	Options   Options  `json:"options,omitempty"` // Additional options.
	KeepAlive string   `json:"keep_alive,omitempty"`
	Stream    bool     `json:"stream,omitempty"`
}

// =========================
// == Model Management API ==
// =========================

// ModelDetails contains specific metadata about a model.
type ModelDetails struct {
	Format        string `json:"format"`
	Family        string `json:"family"`
	ParameterSize string `json:"parameter_size"`
	Quantization  string `json:"quantization_level"`
}

// =========================
// == Version API ==
// =========================

// ✅ **ShowModelRequest**: Used for retrieving model info.
type ShowModelRequest struct {
	Model string `json:"model"`
}

// ✅ **ModelManagementRequest**: Used for creating a model.
type ModelManagementRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner,omitempty"`
}

// CompletionRequest represents a request to generate text completion.
type CompletionRequest struct {
	Model     string   `json:"model"`                // Required: The model name to use.
	Prompt    string   `json:"prompt,omitempty"`     // The prompt to generate a response for.
	Suffix    string   `json:"suffix,omitempty"`     // The text to append after the model's response.
	Images    []string `json:"images,omitempty"`     // List of base64-encoded images for multimodal models.
	Options   Options  `json:"options,omitempty"`    // Advanced model parameters.
	Stream    bool     `json:"stream,omitempty"`     // If true, returns a stream of responses.
	Raw       bool     `json:"raw,omitempty"`        // If true, returns raw model output.
	KeepAlive string   `json:"keep_alive,omitempty"` // Duration to keep the model loaded in memory.
}
