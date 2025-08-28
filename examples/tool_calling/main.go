package main

import (
	"fmt"
	"log"

	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/SamyRai/ollama-go/structures"
)

// getWeather is a simple function that returns the weather for a given location.
func getWeather(location string) (string, error) {
	// In a real application, this would call a weather API.
	// For this example, we'll just return a hardcoded value.
	if location == "Paris" {
		return "The weather in Paris is sunny.", nil
	}
	return fmt.Sprintf("I don't know the weather in %s.", location), nil
}

func main() {
	// Create a new client with default configuration
	cli := client.NewClient(config.DefaultConfig())

	// Define the weather tool
	weatherTool := structures.Tool{
		Type: "function",
		Function: structures.ToolFunction{
			Name:        "get_weather",
			Description: "Get the current weather in a given location.",
			Parameters: map[string]structures.ToolParam{
				"location": {
					Type:        "string",
					Description: "The city and state, e.g. San Francisco, CA",
				},
			},
		},
	}

	// Initial user message
	messages := []structures.Message{
		{Role: "user", Content: "What is the weather like in Paris?"},
	}

	// First API call: Ask the model to use the function
	req := structures.ChatRequest{
		Model:    "llama3.1",
		Messages: messages,
		Tools:    []structures.Tool{weatherTool},
	}

	resp, err := cli.Chat(req, nil)
	if err != nil {
		log.Fatal(err)
	}

	messages = append(messages, resp.Message)

	// If the model wants to call a tool, execute it
	if len(resp.Message.ToolCalls) > 0 {
		for _, toolCall := range resp.Message.ToolCalls {
			if toolCall.Function.Name == "get_weather" {
				// Execute the weather tool
				location := toolCall.Function.Arguments["location"].(string)
				weather, err := getWeather(location)
				if err != nil {
					log.Fatal(err)
				}

				// Add the tool's result to the messages
				toolResult := structures.Message{
					Role:       "tool",
					Content:    weather,
					ToolCallID: "", // Ollama doesn't seem to use tool_call_id yet, but we'll keep the field for OpenAI compatibility.
				}
				messages = append(messages, toolResult)
			}
		}

		// Second API call: Get the final response from the model
		finalReq := structures.ChatRequest{
			Model:    "llama3.1",
			Messages: messages,
		}

		finalResp, err := cli.Chat(finalReq, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(finalResp.Message.Content)
	} else {
		// If the model didn't call a tool, just print the response
		fmt.Println(resp.Message.Content)
	}
}
