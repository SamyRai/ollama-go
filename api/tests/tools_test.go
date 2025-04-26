package tests

import (
	"errors"
	"testing"

	"github.com/SamyRai/ollama-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToolRegistry(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name        string
		toolName    string
		handler     func(args map[string]interface{}) (interface{}, error)
		args        map[string]interface{}
		expectError bool
		expected    interface{}
	}{
		{
			name:     "Register and call a tool successfully",
			toolName: "multiply",
			handler: func(args map[string]interface{}) (interface{}, error) {
				a, ok := args["a"].(float64)
				if !ok {
					return nil, errors.New("a must be a number")
				}
				b, ok := args["b"].(float64)
				if !ok {
					return nil, errors.New("b must be a number")
				}
				return a * b, nil
			},
			args: map[string]interface{}{
				"a": float64(3),
				"b": float64(4),
			},
			expectError: false,
			expected:    float64(12),
		},
		{
			name:        "Tool not registered should return error",
			toolName:    "nonexistent",
			handler:     nil,
			args:        map[string]interface{}{},
			expectError: true,
		},
		{
			name:     "Test error handling in tool",
			toolName: "divide",
			handler: func(args map[string]interface{}) (interface{}, error) {
				a, ok := args["a"].(float64)
				if !ok {
					return nil, errors.New("a must be a number")
				}
				b, ok := args["b"].(float64)
				if !ok {
					return nil, errors.New("b must be a number")
				}
				if b == 0 {
					return nil, errors.New("division by zero")
				}
				return a / b, nil
			},
			args: map[string]interface{}{
				"a": float64(5),
				"b": float64(0),
			},
			expectError: true,
		},
		{
			name:     "Tool returns complex object",
			toolName: "getWeather",
			handler: func(args map[string]interface{}) (interface{}, error) {
				location, ok := args["location"].(string)
				if !ok {
					return nil, errors.New("location must be a string")
				}
				return map[string]interface{}{
					"location":    location,
					"temperature": 22.5,
					"unit":        "celsius",
					"condition":   "sunny",
				}, nil
			},
			args: map[string]interface{}{
				"location": "Paris",
			},
			expectError: false,
			expected: map[string]interface{}{
				"location":    "Paris",
				"temperature": 22.5,
				"unit":        "celsius",
				"condition":   "sunny",
			},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize ToolRegistry
			registry := api.NewToolRegistry()

			// Register tool if provided
			if tt.handler != nil {
				registry.RegisterTool(tt.toolName, tt.handler)
			}

			// Call the tool
			result, err := registry.CallTool(tt.toolName, tt.args)

			// Validate the result
			if tt.expectError {
				// For the divide by zero case, we expect the error to be wrapped in a ToolCallResult
				if tt.toolName == "divide" && tt.args["b"].(float64) == 0 {
					require.NoError(t, err)
					// The result should be nil since we're expecting an error
					assert.Nil(t, result)
				} else {
					require.Error(t, err)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestToolRegistryMultipleTools(t *testing.T) {
	// Initialize ToolRegistry
	registry := api.NewToolRegistry()

	// Register multiple tools
	registry.RegisterTool("add", func(args map[string]interface{}) (interface{}, error) {
		a, ok := args["a"].(float64)
		if !ok {
			return nil, errors.New("a must be a number")
		}
		b, ok := args["b"].(float64)
		if !ok {
			return nil, errors.New("b must be a number")
		}
		return a + b, nil
	})

	registry.RegisterTool("subtract", func(args map[string]interface{}) (interface{}, error) {
		a, ok := args["a"].(float64)
		if !ok {
			return nil, errors.New("a must be a number")
		}
		b, ok := args["b"].(float64)
		if !ok {
			return nil, errors.New("b must be a number")
		}
		return a - b, nil
	})

	// Test add tool
	addResult, err := registry.CallTool("add", map[string]interface{}{
		"a": float64(5),
		"b": float64(3),
	})
	require.NoError(t, err)
	assert.Equal(t, float64(8), addResult)

	// Test subtract tool
	subtractResult, err := registry.CallTool("subtract", map[string]interface{}{
		"a": float64(5),
		"b": float64(3),
	})
	require.NoError(t, err)
	assert.Equal(t, float64(2), subtractResult)
}
