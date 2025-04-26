# Ollama Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/SamyRai/ollama-go.svg)](https://pkg.go.dev/github.com/SamyRai/ollama-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/SamyRai/ollama-go)](https://goreportcard.com/report/github.com/SamyRai/ollama-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go client library for the [Ollama](https://ollama.ai/) API, providing a clean, idiomatic interface for working with local LLMs.

## Features

- Complete API coverage for Ollama's functionality
- Support for both streaming and non-streaming responses
- Chat, completion, and embeddings generation
- Model management (list, create, delete, pull, push)
- Function calling with tool registry
- Context support for cancellation and timeouts
- Comprehensive error handling

## Installation

```bash
go get github.com/SamyRai/ollama-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SamyRai/ollama-go"
)

func main() {
	// Create a new client with default configuration
	client := ollama.New()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Simple chat example
	resp, err := client.Chat().
		WithModel("llama3").
		WithMessage("user", "What is AI?").
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response:", resp.Message.Content)
}
```

## Usage Examples

### Chat Completion

```go
// Create a client with custom configuration
client := ollama.New().
	WithBaseURL("http://localhost:11434").
	WithTimeout(60 * time.Second)

// Build a chat request
resp, err := client.Chat().
	WithModel("llama3").
	WithSystemMessage("You are a helpful assistant.").
	WithMessage("user", "What is the capital of France?").
	WithTemperature(0.7).
	Execute(context.Background())

if err != nil {
	log.Fatalf("Error: %v", err)
}

fmt.Println(resp.Message.Content)
```

### Streaming Chat Completion

```go
err := client.Chat().
	WithModel("llama3").
	WithMessage("user", "Write a short poem about Go programming.").
	Stream(context.Background(), func(resp *ollama.ChatResponse) {
		fmt.Print(resp.Message.Content)
	})

if err != nil {
	log.Fatalf("Error: %v", err)
}
```

### Text Completion

```go
resp, err := client.Completion().
	WithModel("llama3").
	WithPrompt("Once upon a time").
	WithTemperature(0.8).
	Execute(context.Background())

if err != nil {
	log.Fatalf("Error: %v", err)
}

fmt.Println(resp.Response)
```

### Embeddings Generation

```go
resp, err := client.Embeddings().
	WithModel("llama3").
	WithInput("The quick brown fox jumps over the lazy dog").
	Execute(context.Background())

if err != nil {
	log.Fatalf("Error: %v", err)
}

fmt.Printf("Embedding dimensions: %d\n", len(resp.Embeddings[0]))
```

### Model Management

```go
// List available models
models, err := client.Models().List(context.Background())
if err != nil {
	log.Fatalf("Error: %v", err)
}

for _, model := range models.Models {
	fmt.Printf("Model: %s, Size: %d bytes\n", model.Name, model.Size)
}

// Pull a model
err = client.Models().Pull(context.Background(), "llama3")
if err != nil {
	log.Fatalf("Error: %v", err)
}
```

### Function Calling with Tools

```go
// Create a tool registry
registry := ollama.NewToolRegistry()

// Register a weather tool
registry.RegisterTool("getWeather", func(args map[string]interface{}) (interface{}, error) {
	location, ok := args["location"].(string)
	if !ok {
		return nil, fmt.Errorf("location must be a string")
	}

	// In a real application, you would call a weather API here
	return map[string]interface{}{
		"location":    location,
		"temperature": 22,
		"unit":        "celsius",
		"condition":   "sunny",
	}, nil
})

// Define the tool for the model
weatherTool := ollama.Tool{
	Type: "function",
	Function: ollama.ToolFunction{
		Name:        "getWeather",
		Description: "Get the current weather for a location",
		Parameters: map[string]ollama.ToolParam{
			"location": {
				Type:        "string",
				Description: "The city and state, e.g. San Francisco, CA",
			},
		},
	},
}

// Use the tool in a chat
resp, err := client.Chat().
	WithModel("llama3").
	WithMessage("user", "What's the weather like in Paris?").
	WithTools(weatherTool).
	Execute(context.Background())

if err != nil {
	log.Fatalf("Error: %v", err)
}

// Check if the model wants to call a tool
if len(resp.Message.ToolCalls) > 0 {
	toolCall := resp.Message.ToolCalls[0]

	// Execute the tool
	result, err := registry.CallTool(toolCall.Function.Name, toolCall.Function.Arguments)
	if err != nil {
		log.Fatalf("Error calling tool: %v", err)
	}

	// Continue the conversation with the tool result
	resp, err = client.Chat().
		WithModel("llama3").
		WithMessage("user", "What's the weather like in Paris?").
		WithMessage("assistant", resp.Message.Content).
		WithToolResult(toolCall.Function.Name, result).
		Execute(context.Background())

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(resp.Message.Content)
}
```

## Documentation

For complete API documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/SamyRai/ollama-go).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
