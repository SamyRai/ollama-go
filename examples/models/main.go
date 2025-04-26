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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// List available models
	fmt.Println("Listing available models...")
	models, err := client.Models().List(ctx)
	if err != nil {
		log.Fatalf("Error listing models: %v", err)
	}

	fmt.Printf("Found %d models:\n", len(models.Models))
	for _, model := range models.Models {
		fmt.Printf("- %s (%.2f MB)\n", model.Name, float64(model.Size)/(1024*1024))
	}

	// Get details about a specific model
	if len(models.Models) > 0 {
		modelName := models.Models[0].Name
		fmt.Printf("\nGetting details for model: %s\n", modelName)
		details, err := client.Models().Show(ctx, modelName)
		if err != nil {
			log.Printf("Error getting model details: %v", err)
		} else {
			fmt.Printf("Model: %s\n", details.Name)
			fmt.Printf("Version: %s\n", details.Version)
			fmt.Printf("Description: %s\n", details.Description)
			fmt.Printf("Tags: %v\n", details.Tags)
		}
	}

	// Example of pulling a model (commented out to avoid actual download)
	/*
		fmt.Println("\nPulling a model (this may take a while)...")
		err = client.Models().Pull(ctx, "llama3:latest")
		if err != nil {
			log.Fatalf("Error pulling model: %v", err)
		}
		fmt.Println("Model pulled successfully!")
	*/

	// Get running processes
	fmt.Println("\nGetting running processes...")
	processes, err := client.Status().GetRunningProcesses(ctx)
	if err != nil {
		log.Printf("Error getting running processes: %v", err)
	} else {
		fmt.Printf("Found %d running models:\n", len(processes.Models))
		for _, proc := range processes.Models {
			fmt.Printf("- %s (%.2f MB VRAM)\n", proc.Model, float64(proc.VRAMSize)/(1024*1024))
		}
	}

	// Get API version
	fmt.Println("\nGetting API version...")
	version, err := client.Status().GetVersion(ctx)
	if err != nil {
		log.Printf("Error getting API version: %v", err)
	} else {
		fmt.Printf("Ollama API version: %s\n", version.Version)
	}
}
