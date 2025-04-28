// Package main demonstrates how to use the Ollama Go library with function calling tools.
// This example shows how to register, define, and execute tool functions in a chat conversation.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SamyRai/ollama-go"
)

// WeatherData represents weather information for a specific location.
type WeatherData struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Condition   string  `json:"condition"`
}

func main() {
	// Create a new client with default configuration
	client := ollama.New()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a tool registry
	registry := ollama.NewToolRegistry()

	// Register a weather tool
	registry.RegisterTool("getWeather", func(args map[string]interface{}) (interface{}, error) {
		// Extract the location from the arguments
		location, ok := args["location"].(string)
		if !ok {
			return nil, fmt.Errorf("location must be a string")
		}

		// In a real application, you would call a weather API here
		// For this example, we'll return mock data
		weatherData := WeatherData{
			Location:    location,
			Temperature: 22.5,
			Unit:        "celsius",
			Condition:   "sunny",
		}

		return weatherData, nil
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
	fmt.Println("Sending chat request with tool...")
	resp, err := client.Chat().
		WithModel("llama3").
		WithSystemMessage("You are a helpful assistant that can provide weather information.").
		WithMessage("user", "What's the weather like in Paris?").
		WithTools(weatherTool).
		Execute(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response:", resp.Message.Content)

	// Check if the model wants to call a tool
	if len(resp.Message.ToolCalls) > 0 {
		fmt.Println("\nModel requested to call a tool:")
		toolCall := resp.Message.ToolCalls[0]
		fmt.Printf("Tool: %s\n", toolCall.Function.Name)

		// Pretty print the arguments
		argsJSON, _ := json.MarshalIndent(toolCall.Function.Arguments, "", "  ")
		fmt.Printf("Arguments: %s\n", argsJSON)

		// Execute the tool
		result, err := registry.CallTool(toolCall.Function.Name, toolCall.Function.Arguments)
		if err != nil {
			log.Fatalf("Error calling tool: %v", err)
		}

		// Pretty print the result
		resultJSON, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("\nTool result: %s\n", resultJSON)

		// Continue the conversation with the tool result
		fmt.Println("\nContinuing conversation with tool result...")

		// Convert result to string for the tool message
		resultStr, _ := json.Marshal(result)

		finalResp, err := client.Chat().
			WithModel("llama3").
			WithSystemMessage("You are a helpful assistant that can provide weather information.").
			WithMessage("user", "What's the weather like in Paris?").
			WithMessage("assistant", resp.Message.Content).
			WithMessage("tool", string(resultStr)).
			Execute(ctx)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		fmt.Println("Final response:", finalResp.Message.Content)
	}
}
