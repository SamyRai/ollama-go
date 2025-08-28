# Ollama-Go Client

A simple and easy-to-use Go client for the Ollama API.

This library provides a convenient way to interact with the Ollama API, allowing you to easily integrate Ollama's powerful language models into your Go applications.

## Features

-   Chat with Ollama models
-   Generate text completions
-   Manage models (create, delete, copy, pull, push)
-   Support for both streaming and non-streaming responses

## Installation

To install the Ollama-Go client, use `go get`:

```bash
go get github.com/SamyRai/ollama-go
```

## Usage

Here's a simple example of how to use the client to chat with an Ollama model:

```go
package main

import (
	"fmt"
	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/SamyRai/ollama-go/structures"
)

func main() {
	// Create a new client with default configuration
	cli := client.NewClient(config.DefaultConfig())

	// Define the chat request
	req := structures.ChatRequest{
		Model: "llama3.1",
		Messages: []structures.Message{
			{Role: "user", Content: "What is AI?"},
		},
	}

	// Send the request
	resp, err := cli.Chat(req, nil)
	if err != nil {
		panic(err)
	}

	// Print the response
	fmt.Println(resp.Message.Content)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any feedback or suggestions.
