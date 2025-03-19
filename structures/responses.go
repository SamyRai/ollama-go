package structures

import "time"

// CompletionResponse represents the model's response to a text generation request.
type CompletionResponse struct {
	Model              string                 `json:"model"`                 // Name of the model used.
	CreatedAt          time.Time              `json:"created_at"`            // Timestamp of response creation.
	Response           string                 `json:"response"`              // Generated text.
	Done               bool                   `json:"done"`                  // Whether the generation is complete.
	DoneReason         string                 `json:"done_reason,omitempty"` // Reason for stopping (if applicable).
	Context            []int                  `json:"context,omitempty"`     // Optional: Encoding of conversation for memory.
	TotalDuration      int64                  `json:"total_duration"`        // Total time spent on generation (ns).
	LoadDuration       int64                  `json:"load_duration"`         // Time spent loading the model (ns).
	PromptEvalCount    int                    `json:"prompt_eval_count"`     // Number of tokens in the prompt.
	PromptEvalDuration int64                  `json:"prompt_eval_duration"`  // Time spent evaluating prompt (ns).
	EvalCount          int                    `json:"eval_count"`            // Number of tokens generated.
	EvalDuration       int64                  `json:"eval_duration"`         // Time spent generating tokens (ns).
	Metadata           map[string]interface{} `json:"metadata,omitempty"`    // Optional: Additional metadata.
}

// =========================
// == Chat API ==
// =========================

// ChatResponse represents the model's reply in a chat conversation.
type ChatResponse struct {
	Model      string                 `json:"model"`                 // Model used for the response.
	CreatedAt  time.Time              `json:"created_at"`            // Timestamp of response creation.
	Message    Message                `json:"message"`               // Assistant's message.
	DoneReason string                 `json:"done_reason,omitempty"` // Optional: Reason for stopping.
	Done       bool                   `json:"done"`                  // Whether the chat response is complete.
	ToolCalls  []ToolCall             `json:"tool_calls,omitempty"`  // Tool calls (structured return).
	Metadata   map[string]interface{} `json:"metadata,omitempty"`    // Optional: Additional metadata.
}

// =========================
// == Embeddings API ==
// =========================

// EmbeddingResponse contains generated embeddings.
type EmbeddingResponse struct {
	Model      string      `json:"model"`      // Model name.
	Embeddings [][]float32 `json:"embeddings"` // Embedding vectors.
}

// =========================
// == Model Management API ==
// =========================

// ModelListResponse contains available local models.
type ModelListResponse struct {
	Models []ModelInfo `json:"models"`
}

// ModelInfo contains details about a local model.
type ModelInfo struct {
	Name       string       `json:"name"`
	ModifiedAt time.Time    `json:"modified_at"`
	Size       int64        `json:"size"`
	Digest     string       `json:"digest"`
	Details    ModelDetails `json:"details"`
}

// =========================
// == Model Process API ==
// =========================

// ModelProcessResponse lists running models.
type ModelProcessResponse struct {
	Models []ModelProcess `json:"models"`
}

type ModelProcess struct {
	Name      string       `json:"name"`
	Model     string       `json:"model"`
	Size      int64        `json:"size"`
	Digest    string       `json:"digest"`
	ExpiresAt time.Time    `json:"expires_at"`
	VRAMSize  int64        `json:"vram_size"`
	Details   ModelDetails `json:"details"`
}

// =========================
// == Version API ==
// =========================

// VersionResponse returns the server version.
type VersionResponse struct {
	Version string `json:"version"`
}

// ✅ **ShowModelResponse**: Contains model details.
type ShowModelResponse struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
