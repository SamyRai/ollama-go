// Package main demonstrates how to use the Ollama Go library for generating embeddings.
// This example shows how to generate embeddings for both single and multiple text inputs.
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

	// Generate embeddings for a text
	fmt.Println("Generating embeddings...")
	resp, err := client.Embeddings().
		WithModel("llama3").
		WithInput("The quick brown fox jumps over the lazy dog").
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the first few dimensions of the embedding
	fmt.Printf("Model: %s\n", resp.Model)
	fmt.Printf("Embedding dimensions: %d\n", len(resp.Embeddings[0]))
	fmt.Printf("First 5 dimensions: %v\n", resp.Embeddings[0][:5])

	// Generate embeddings for multiple texts
	fmt.Println("\nGenerating embeddings for multiple texts...")
	multiResp, err := client.Embeddings().
		WithModel("llama3").
		WithInputs([]string{
			"The quick brown fox jumps over the lazy dog",
			"The five boxing wizards jump quickly",
		}).
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Number of embeddings: %d\n", len(multiResp.Embeddings))
	fmt.Printf("Dimensions of first embedding: %d\n", len(multiResp.Embeddings[0]))
	fmt.Printf("Dimensions of second embedding: %d\n", len(multiResp.Embeddings[1]))
}
