package testutils

import (
	"github.com/SamyRai/ollama-go/internal/structures"
)

// Constants for test data
const (
	TestModel = "llama3.1"
)

// CreateChatRequest creates a simple chat request for testing.
func CreateChatRequest(userMessage string) structures.ChatRequest {
	return structures.ChatRequest{
		Model: TestModel,
		Messages: []structures.Message{
			{Role: "user", Content: userMessage},
		},
	}
}

// CreateChatRequestWithSystemMessage creates a chat request with a system message for testing.
func CreateChatRequestWithSystemMessage(systemMessage, userMessage string) structures.ChatRequest {
	return structures.ChatRequest{
		Model: TestModel,
		Messages: []structures.Message{
			{Role: "system", Content: systemMessage},
			{Role: "user", Content: userMessage},
		},
	}
}

// CreateChatRequestWithTools creates a chat request with tools for testing.
func CreateChatRequestWithTools(userMessage string, tools []structures.Tool) structures.ChatRequest {
	return structures.ChatRequest{
		Model: TestModel,
		Messages: []structures.Message{
			{Role: "user", Content: userMessage},
		},
		Tools: tools,
	}
}

// CreateCompletionRequest creates a simple completion request for testing.
func CreateCompletionRequest(prompt string) structures.CompletionRequest {
	return structures.CompletionRequest{
		Model:  TestModel,
		Prompt: prompt,
	}
}

// CreateEmbeddingRequest creates a simple embedding request for testing.
func CreateEmbeddingRequest(input string) structures.EmbeddingRequest {
	return structures.EmbeddingRequest{
		Model: TestModel,
		Input: []string{input},
	}
}

// CreateWeatherTool creates a weather tool for testing.
func CreateWeatherTool() structures.Tool {
	return structures.Tool{
		Type: "function",
		Function: structures.ToolFunction{
			Name:        "getWeather",
			Description: "Retrieve the weather",
			Parameters: map[string]structures.ToolParam{
				"location": {
					Type:        "string",
					Description: "The location for the weather forecast",
				},
			},
		},
	}
}

// CreateCalculatorTool creates a calculator tool for testing.
func CreateCalculatorTool() structures.Tool {
	return structures.Tool{
		Type: "function",
		Function: structures.ToolFunction{
			Name:        "calculate",
			Description: "Perform a calculation",
			Parameters: map[string]structures.ToolParam{
				"operation": {
					Type:        "string",
					Description: "The operation to perform (add, subtract, multiply, divide)",
					Enum:        []string{"add", "subtract", "multiply", "divide"},
				},
				"a": {
					Type:        "number",
					Description: "The first number",
				},
				"b": {
					Type:        "number",
					Description: "The second number",
				},
			},
		},
	}
}
