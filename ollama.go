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
func (c *Client) Chat() *api.ChatBuilder {
	return api.NewChatBuilder(c.client)
}

// Completion returns a new CompletionBuilder for constructing completion requests.
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

// WithTemperature sets the temperature parameter.
func WithTemperature(temperature float64) func(*Options) {
	return api.WithTemperature(temperature)
}

// WithTopP sets the top_p parameter.
func WithTopP(topP float64) func(*Options) {
	return api.WithTopP(topP)
}

// WithTopK sets the top_k parameter.
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

// WithRepeatPenalty sets the repeat_penalty parameter.
func WithRepeatPenalty(repeatPenalty float64) func(*Options) {
	return api.WithRepeatPenalty(repeatPenalty)
}

// WithRepeatLastN sets the repeat_last_n parameter.
func WithRepeatLastN(repeatLastN int) func(*Options) {
	return api.WithRepeatLastN(repeatLastN)
}

// WithFrequencyPenalty sets the frequency_penalty parameter.
func WithFrequencyPenalty(frequencyPenalty float64) func(*Options) {
	return api.WithFrequencyPenalty(frequencyPenalty)
}

// WithPresencePenalty sets the presence_penalty parameter.
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
type (
	// Chat API
	ChatResponse     = api.ChatResponse
	Message          = api.Message
	Tool             = api.Tool
	ToolFunction     = api.ToolFunction
	ToolParam        = api.ToolParam
	ToolCall         = api.ToolCall
	ToolCallFunction = api.ToolCallFunction

	// Completion API
	CompletionResponse = api.CompletionResponse

	// Embeddings API
	EmbeddingResponse = api.EmbeddingResponse

	// Model Management API
	ModelListResponse = api.ModelListResponse
	ModelInfo         = api.ModelInfo
	ShowModelResponse = api.ShowModelResponse

	// Status API
	VersionResponse      = api.VersionResponse
	ModelProcessResponse = api.ModelProcessResponse
	ModelProcess         = api.ModelProcess
)

// RawClient returns the underlying OllamaClient for advanced usage.
// This is provided for cases where the high-level API doesn't provide
// the needed functionality.
func (c *Client) RawClient() *client.OllamaClient {
	return c.client
}
