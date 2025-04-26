package api

import (
	"context"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/structures"
)

// ChatResponse represents the model's reply in a chat conversation.
type ChatResponse = structures.ChatResponse

// Message represents a single message in a chat conversation.
type Message = structures.Message

// Tool defines a callable function available to models.
type Tool = structures.Tool

// ToolFunction describes the function a model can invoke.
type ToolFunction = structures.ToolFunction

// ToolParam defines a function parameter.
type ToolParam = structures.ToolParam

// ToolCall represents an invocation of a tool.
type ToolCall = structures.ToolCall

// ToolCallFunction contains function call arguments.
type ToolCallFunction = structures.ToolCallFunction

// ChatBuilder provides a fluent interface for building chat requests.
type ChatBuilder struct {
	client   *client.OllamaClient
	model    string
	messages []Message
	tools    []Tool
	format   string
	options  structures.Options
	stream   bool
}

// NewChatBuilder creates a new ChatBuilder.
func NewChatBuilder(client *client.OllamaClient) *ChatBuilder {
	return &ChatBuilder{
		client:   client,
		messages: make([]Message, 0),
		tools:    make([]Tool, 0),
	}
}

// WithModel sets the model to use for the chat.
func (b *ChatBuilder) WithModel(model string) *ChatBuilder {
	b.model = model
	return b
}

// WithSystemMessage adds a system message to the chat.
func (b *ChatBuilder) WithSystemMessage(content string) *ChatBuilder {
	b.messages = append(b.messages, Message{
		Role:    "system",
		Content: content,
	})
	return b
}

// WithMessage adds a message to the chat.
func (b *ChatBuilder) WithMessage(role, content string) *ChatBuilder {
	b.messages = append(b.messages, Message{
		Role:    role,
		Content: content,
	})
	return b
}

// WithMessages sets all messages for the chat.
func (b *ChatBuilder) WithMessages(messages []Message) *ChatBuilder {
	b.messages = messages
	return b
}

// WithTools adds tools to the chat.
func (b *ChatBuilder) WithTools(tools ...Tool) *ChatBuilder {
	b.tools = append(b.tools, tools...)
	return b
}

// WithFormat sets the response format.
func (b *ChatBuilder) WithFormat(format string) *ChatBuilder {
	b.format = format
	return b
}

// WithTemperature sets the temperature for the chat.
func (b *ChatBuilder) WithTemperature(temperature float64) *ChatBuilder {
	b.options.Temperature = temperature
	return b
}

// WithTopP sets the top_p value for the chat.
func (b *ChatBuilder) WithTopP(topP float64) *ChatBuilder {
	b.options.TopP = topP
	return b
}

// WithTopK sets the top_k value for the chat.
func (b *ChatBuilder) WithTopK(topK int) *ChatBuilder {
	b.options.TopK = topK
	return b
}

// WithOptions sets all options for the chat.
func (b *ChatBuilder) WithOptions(options structures.Options) *ChatBuilder {
	b.options = options
	return b
}

// WithToolResult adds a tool result message to the chat.
func (b *ChatBuilder) WithToolResult(toolName string, result interface{}) *ChatBuilder {
	// Add the tool result as a message from the tool
	b.messages = append(b.messages, Message{
		Role:    "tool",
		Content: toolName + " result: " + result.(string),
	})
	return b
}

// Execute sends the chat request and returns the response.
func (b *ChatBuilder) Execute(ctx context.Context) (*ChatResponse, error) {
	req := structures.ChatRequest{
		Model:    b.model,
		Messages: b.messages,
		Tools:    b.tools,
		Format:   b.format,
		Options:  b.options,
		Stream:   false,
	}

	// Use a no-op callback since we're not streaming
	resp, err := b.client.Chat(req, func(resp structures.ChatResponse) {})
	return resp, err
}

// Stream sends the chat request and streams the response through the callback.
func (b *ChatBuilder) Stream(ctx context.Context, callback func(*ChatResponse)) error {
	req := structures.ChatRequest{
		Model:    b.model,
		Messages: b.messages,
		Tools:    b.tools,
		Format:   b.format,
		Options:  b.options,
		Stream:   true,
	}

	_, err := b.client.Chat(req, func(resp structures.ChatResponse) {
		callback(&resp)
	})
	return err
}
