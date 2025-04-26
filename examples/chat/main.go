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
	fmt.Println("Sending chat request...")
	resp, err := client.Chat().
		WithModel("llama3").
		WithSystemMessage("You are a helpful assistant.").
		WithMessage("user", "What is artificial intelligence in 3 sentences?").
		WithTemperature(0.7).
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response:", resp.Message.Content)

	// Example with streaming
	fmt.Println("\nStreaming response...")
	err = client.Chat().
		WithModel("llama3").
		WithMessage("user", "Write a short poem about Go programming.").
		Stream(ctx, func(resp *ollama.ChatResponse) {
			fmt.Print(resp.Message.Content)
		})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("\nDone!")
}
