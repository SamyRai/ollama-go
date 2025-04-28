// Package structures provides data structures used throughout the Ollama Go client.
// It includes request and response objects, options, and other types.
package structures

// Options defines customizable parameters for model behavior.
type Options struct {
	Temperature      float64 `json:"temperature,omitempty"`       // Controls creativity vs. coherence.
	TopP             float64 `json:"top_p,omitempty"`             // Nucleus sampling parameter.
	TopK             int     `json:"top_k,omitempty"`             // Limits highest probability tokens.
	Mirostat         int     `json:"mirostat,omitempty"`          // Mirostat sampling mode.
	MirostatTau      float64 `json:"mirostat_tau,omitempty"`      // Target surprise value for Mirostat.
	MirostatEta      float64 `json:"mirostat_eta,omitempty"`      // Learning rate for Mirostat.
	RepeatPenalty    float64 `json:"repeat_penalty,omitempty"`    // Penalizes repeated tokens.
	RepeatLastN      int     `json:"repeat_last_n,omitempty"`     // Tokens considered for repetition penalty.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"` // Penalizes frequent tokens.
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`  // Penalizes existing tokens in context.
	TFS              float64 `json:"tfs,omitempty"`               // Tail Free Sampling parameter.
	TopA             float64 `json:"top_a,omitempty"`             // Alternative sampling parameter.
	TypicalP         float64 `json:"typical_p,omitempty"`         // Typical probability threshold.
	Grammar          string  `json:"grammar,omitempty"`           // Enforces specific grammar on output.
	// Additional parameters may be added here as needed.
}
