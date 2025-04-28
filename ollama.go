// Package ollama provides a comprehensive Go client for the Ollama API.
//
// This package allows Go applications to interact with Ollama's local LLM capabilities,
// including chat, completion, embeddings, and model management.
package ollama

import (
	"io"
	"net/http"
	"time"

	"github.com/SamyRai/ollama-go/api"
	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/SamyRai/ollama-go/internal/utils"
)

// Version returns the current version of the ollama-go library.
func Version() string {
	return "ollama-go v1.0.0"
}

// Client is the main interface for interacting with the Ollama API.
type Client struct {
	client *client.OllamaClient
	config *config.Config
}

// New creates a new Ollama client with default configuration.
func New() *Client {
	cfg := config.DefaultConfig()
	return &Client{
		client: client.NewClient(cfg),
		config: cfg,
	}
}

// WithBaseURL sets a custom base URL for the Ollama API.
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.config.BaseURL = baseURL
	c.client = client.NewClient(c.config)
	return c
}

// WithTimeout sets a custom timeout for API requests.
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.config.Timeout = timeout
	c.client = client.NewClient(c.config)
	return c
}

// WithHTTPClient sets a custom HTTP client.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.client.HTTPClient = httpClient
	return c
}

// WithAPIKey sets an API key for authentication (if required).
func (c *Client) WithAPIKey(apiKey string) *Client {
	c.config.APIKey = apiKey
	c.client = client.NewClient(c.config)
	return c
}

// WithDebugMode enables or disables debug logging.
func (c *Client) WithDebugMode(debug bool) *Client {
	c.config.Debug = debug
	c.client = client.NewClient(c.config)
	return c
}

// SetLogLevel sets the logging level for the client.
// Valid levels are: "NONE", "ERROR", "WARN", "INFO", "DEBUG"
func SetLogLevel(level string) {
	var logLevel utils.LogLevel

	switch level {
	case "NONE":
		logLevel = utils.LogLevelNone
	case "ERROR":
		logLevel = utils.LogLevelError
	case "WARN":
		logLevel = utils.LogLevelWarn
	case "INFO":
		logLevel = utils.LogLevelInfo
	case "DEBUG":
		logLevel = utils.LogLevelDebug
	default:
		logLevel = utils.LogLevelInfo // Default to INFO if invalid level
	}

	utils.SetGlobalLevel(logLevel)
}

// SetLogOutput sets the output destination for logs.
// This can be used to redirect logs to a file or other writer.
func SetLogOutput(w io.Writer) {
	utils.SetGlobalOutput(w)
}

// Chat returns a new ChatBuilder for constructing chat requests.
//
// The ChatBuilder provides a fluent interface for creating chat requests:
//
//	client.Chat().
//		WithModel("llama3").
//		WithSystemMessage("You are a helpful assistant.").
//		WithMessage("user", "What is artificial intelligence?").
//		Execute(ctx)
//
// For streaming responses, use the Stream method instead of Execute:
//
//	client.Chat().
//		WithModel("llama3").
//		WithMessage("user", "Write a poem.").
//		Stream(ctx, func(resp *ChatResponse) {
//			fmt.Print(resp.Message.Content)
//		})
func (c *Client) Chat() *api.ChatBuilder {
	return api.NewChatBuilder(c.client)
}

// Completion returns a new CompletionBuilder for constructing completion requests.
//
// The CompletionBuilder provides a fluent interface for text generation:
//
//	client.Completion().
//		WithModel("llama3").
//		WithPrompt("Once upon a time").
//		WithTemperature(0.8).
//		Execute(ctx)
//
// For streaming completions, use the Stream method instead of Execute:
//
//	client.Completion().
//		WithModel("llama3").
//		WithPrompt("Write a story about").
//		Stream(ctx, func(resp *CompletionResponse) {
//			fmt.Print(resp.Response)
//		})
func (c *Client) Completion() *api.CompletionBuilder {
	return api.NewCompletionBuilder(c.client)
}

// Embeddings returns a new EmbeddingsBuilder for constructing embeddings requests.
func (c *Client) Embeddings() *api.EmbeddingsBuilder {
	return api.NewEmbeddingsBuilder(c.client)
}

// Models returns a ModelManager for model-related operations.
func (c *Client) Models() *api.ModelManager {
	return api.NewModelManager(c.client)
}

// Status returns a StatusManager for status-related operations.
func (c *Client) Status() *api.StatusManager {
	return api.NewStatusManager(c.client)
}

// NewToolRegistry creates a new tool registry.
func NewToolRegistry() *api.ToolRegistry {
	return api.NewToolRegistry()
}

// Options defines customizable parameters for model behavior.
type Options = api.Options

// NewOptions creates a new Options instance with default values.
func NewOptions() *Options {
	return api.NewOptions()
}

// ApplyOptions applies a list of option functions to an Options instance.
func ApplyOptions(o *Options, opts ...func(*Options)) {
	api.ApplyOptions(o, opts...)
}

// WithTemperature sets the temperature parameter for text generation.
//
// Temperature controls randomness in generation:
//   - Values closer to 0 produce more deterministic responses
//   - Values closer to 1 produce more creative and diverse responses
//   - The recommended range is 0.0 - 1.0, with 0.7 being a common default
func WithTemperature(temperature float64) func(*Options) {
	return api.WithTemperature(temperature)
}

// WithTopP sets the top_p parameter for nucleus sampling.
//
// Top-p (nucleus) sampling:
//   - Only considers tokens whose cumulative probability exceeds the probability threshold p
//   - Lower values (0.5) are more focused and deterministic
//   - Higher values (0.9) allow more diversity but may be less coherent
//   - A value of 1.0 disables this effect
//   - Often produces more diverse outputs than temperature sampling alone
func WithTopP(topP float64) func(*Options) {
	return api.WithTopP(topP)
}

// WithTopK sets the top_k parameter for limiting vocabulary in sampling.
//
// Top-k sampling:
//   - Limits token selection to the k highest probability tokens
//   - Lower values (10) focus on most likely tokens, making output more conservative
//   - Higher values (50+) allow for more diversity in generation
//   - Works well when combined with top_p and temperature
//   - A value of 0 disables this effect (all tokens considered)
func WithTopK(topK int) func(*Options) {
	return api.WithTopK(topK)
}

// WithMirostat sets the mirostat parameter.
func WithMirostat(mirostat int) func(*Options) {
	return api.WithMirostat(mirostat)
}

// WithMirostatTau sets the mirostat_tau parameter.
func WithMirostatTau(mirostatTau float64) func(*Options) {
	return api.WithMirostatTau(mirostatTau)
}

// WithMirostatEta sets the mirostat_eta parameter.
func WithMirostatEta(mirostatEta float64) func(*Options) {
	return api.WithMirostatEta(mirostatEta)
}

// WithRepeatPenalty sets the repeat_penalty parameter to discourage repetition in generation.
//
// Repeat penalty:
//   - Controls how strongly to penalize repetitions of the same tokens
//   - Higher values (1.1 - 2.0) reduce repetition more aggressively
//   - A value of 1.0 applies no penalty
//   - Particularly useful for creative text generation like stories
//   - Works in conjunction with RepeatLastN to determine how far back to look
func WithRepeatPenalty(repeatPenalty float64) func(*Options) {
	return api.WithRepeatPenalty(repeatPenalty)
}

// WithRepeatLastN sets the repeat_last_n parameter for repetition control.
//
// Repeat last N:
//   - Specifies how many previous tokens to consider for the repeat penalty
//   - Higher values (e.g., 64, 128, 256) reduce repetition over larger contexts
//   - Lower values (e.g., 8, 16) only penalize immediate repetitions
//   - Works in conjunction with RepeatPenalty parameter
//   - Setting this too high may prevent intentional repetitions like list numbering
func WithRepeatLastN(repeatLastN int) func(*Options) {
	return api.WithRepeatLastN(repeatLastN)
}

// WithFrequencyPenalty sets the frequency_penalty parameter to reduce repetition.
//
// Frequency penalty:
//   - Penalizes tokens based on how frequently they've appeared in the generated text
//   - Higher values (0.1 - 1.0) reduce the likelihood of repeating the same words
//   - Positive values favor tokens that appear less often in the output
//   - Useful for encouraging diversity in longer generations
//   - Differs from repeat_penalty by considering overall frequency rather than sequences
func WithFrequencyPenalty(frequencyPenalty float64) func(*Options) {
	return api.WithFrequencyPenalty(frequencyPenalty)
}

// WithPresencePenalty sets the presence_penalty parameter to reduce topic repetition.
//
// Presence penalty:
//   - Penalizes tokens based on their presence in the generated text so far
//   - Higher values (0.1 - 1.0) encourage the model to talk about new topics
//   - A value of 0.0 applies no penalty
//   - Helpful for open-ended conversations to prevent the model from fixating on topics
//   - Unlike frequency penalty, this considers presence (used or not) rather than frequency
func WithPresencePenalty(presencePenalty float64) func(*Options) {
	return api.WithPresencePenalty(presencePenalty)
}

// WithTFS sets the tfs parameter.
func WithTFS(tfs float64) func(*Options) {
	return api.WithTFS(tfs)
}

// WithTopA sets the top_a parameter.
func WithTopA(topA float64) func(*Options) {
	return api.WithTopA(topA)
}

// WithTypicalP sets the typical_p parameter.
func WithTypicalP(typicalP float64) func(*Options) {
	return api.WithTypicalP(typicalP)
}

// WithGrammar sets the grammar parameter.
func WithGrammar(grammar string) func(*Options) {
	return api.WithGrammar(grammar)
}

// Type aliases for public API

// Message represents a single message in a chat conversation.
type Message = api.Message

// ChatResponse represents the model's reply in a chat conversation.
type ChatResponse = api.ChatResponse

// Tool defines a callable function available to models.
type Tool = api.Tool

// ToolFunction describes the function a model can invoke.
type ToolFunction = api.ToolFunction

// ToolParam defines a function parameter.
type ToolParam = api.ToolParam

// ToolCall represents an invocation of a tool.
type ToolCall = api.ToolCall

// ToolCallFunction contains function call arguments.
type ToolCallFunction = api.ToolCallFunction

// CompletionResponse represents the model's response to a text generation request.
type CompletionResponse = api.CompletionResponse

// EmbeddingResponse contains generated embeddings.
type EmbeddingResponse = api.EmbeddingResponse

// ModelListResponse contains available local models.
type ModelListResponse = api.ModelListResponse

// ModelInfo contains details about a local model.
type ModelInfo = api.ModelInfo

// ShowModelResponse contains model details.
type ShowModelResponse = api.ShowModelResponse

// VersionResponse returns the server version.
type VersionResponse = api.VersionResponse

// ModelProcessResponse lists running models.
type ModelProcessResponse = api.ModelProcessResponse

// ModelProcess contains information about a running model process in Ollama.
type ModelProcess = api.ModelProcess

// RawClient returns the underlying OllamaClient for advanced usage.
// This is provided for cases where the high-level API doesn't provide
// the needed functionality.
func (c *Client) RawClient() *client.OllamaClient {
	return c.client
}
