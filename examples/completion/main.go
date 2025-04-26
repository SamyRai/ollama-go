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

	// Simple completion example
	fmt.Println("Sending completion request...")
	resp, err := client.Completion().
		WithModel("llama3").
		WithPrompt("Once upon a time in a land far away").
		WithTemperature(0.8).
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response:", resp.Response)

	// Example with streaming
	fmt.Println("\nStreaming response...")
	err = client.Completion().
		WithModel("llama3").
		WithPrompt("Write a short story about a robot learning to paint").
		WithTemperature(0.7).
		Stream(ctx, func(resp *ollama.CompletionResponse) {
			fmt.Print(resp.Response)
		})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("\nDone!")

	// Example with more options
	fmt.Println("\nCompletion with more options...")

	// Create options with functional options
	options := ollama.NewOptions()
	ollama.ApplyOptions(options,
		ollama.WithTemperature(0.7),
		ollama.WithTopP(0.9),
		ollama.WithTopK(40),
		ollama.WithRepeatPenalty(1.1),
	)

	resp, err = client.Completion().
		WithModel("llama3").
		WithPrompt("List 5 benefits of artificial intelligence:").
		WithOptions(*options).
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response with custom options:", resp.Response)
}
