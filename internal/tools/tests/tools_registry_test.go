package tests

import (
	"errors"
	"testing"

	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/SamyRai/ollama-go/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToolRegistry(t *testing.T) {
	// Test cases using table-driven testing
	tests := []struct {
		name      string
		toolName  string
		toolFunc  func(args structures.ToolCallFunction) (structures.ToolCallResult, error)
		args      structures.ToolCallFunction
		expectErr bool
		expected  structures.ToolCallResult
	}{
		{
			name:     "Register and call a tool successfully",
			toolName: "multiply",
			toolFunc: func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
				a, ok := args.Arguments["a"].(float64)
				if !ok {
					return structures.ToolCallResult{}, errors.New("a must be a number")
				}
				b, ok := args.Arguments["b"].(float64)
				if !ok {
					return structures.ToolCallResult{}, errors.New("b must be a number")
				}
				result := a * b
				return structures.ToolCallResult{
					Status: "success",
					Result: result,
				}, nil
			},
			args: structures.ToolCallFunction{
				Name: "multiply",
				Arguments: map[string]interface{}{
					"a": float64(3),
					"b": float64(4),
				},
			},
			expectErr: false,
			expected: structures.ToolCallResult{
				Status: "success",
				Result: float64(12),
			},
		},
		{
			name:      "Tool not registered should return error",
			toolName:  "nonexistent",
			toolFunc:  nil,
			args:      structures.ToolCallFunction{Name: "nonexistent"},
			expectErr: true,
		},
		{
			name:     "Test error handling in tool",
			toolName: "divide",
			toolFunc: func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
				a, ok := args.Arguments["a"].(float64)
				if !ok {
					return structures.ToolCallResult{}, errors.New("a must be a number")
				}
				b, ok := args.Arguments["b"].(float64)
				if !ok {
					return structures.ToolCallResult{}, errors.New("b must be a number")
				}
				if b == 0 {
					return structures.ToolCallResult{}, errors.New("division by zero")
				}
				result := a / b
				return structures.ToolCallResult{
					Status: "success",
					Result: result,
				}, nil
			},
			args: structures.ToolCallFunction{
				Name: "divide",
				Arguments: map[string]interface{}{
					"a": float64(5),
					"b": float64(0),
				},
			},
			expectErr: true,
		},
		{
			name:     "Tool returns error status",
			toolName: "failingTool",
			toolFunc: func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
				return structures.ToolCallResult{
					Status: "error",
					Error:  "This tool always fails",
				}, nil
			},
			args: structures.ToolCallFunction{
				Name:      "failingTool",
				Arguments: map[string]interface{}{},
			},
			expectErr: false,
			expected: structures.ToolCallResult{
				Status: "error",
				Error:  "This tool always fails",
			},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize ToolRegistry
			registry := tools.NewRegistry()

			// Register tool if provided
			if tt.toolFunc != nil {
				registry.RegisterTool(tt.toolName, tt.toolFunc)
			}

			// Call the tool
			result, err := registry.CallTool(tt.toolName, tt.args)

			// Validate the result
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected.Status, result.Status)
				assert.Equal(t, tt.expected.Result, result.Result)
				if tt.expected.Error != "" {
					assert.Equal(t, tt.expected.Error, result.Error)
				}
			}
		})
	}
}

func TestToolRegistryMultipleTools(t *testing.T) {
	// Initialize ToolRegistry
	registry := tools.NewRegistry()

	// Register multiple tools
	registry.RegisterTool("add", func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
		a, ok := args.Arguments["a"].(float64)
		if !ok {
			return structures.ToolCallResult{}, errors.New("a must be a number")
		}
		b, ok := args.Arguments["b"].(float64)
		if !ok {
			return structures.ToolCallResult{}, errors.New("b must be a number")
		}
		return structures.ToolCallResult{
			Status: "success",
			Result: a + b,
		}, nil
	})

	registry.RegisterTool("subtract", func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
		a, ok := args.Arguments["a"].(float64)
		if !ok {
			return structures.ToolCallResult{}, errors.New("a must be a number")
		}
		b, ok := args.Arguments["b"].(float64)
		if !ok {
			return structures.ToolCallResult{}, errors.New("b must be a number")
		}
		return structures.ToolCallResult{
			Status: "success",
			Result: a - b,
		}, nil
	})

	// Test add tool
	addResult, err := registry.CallTool("add", structures.ToolCallFunction{
		Name: "add",
		Arguments: map[string]interface{}{
			"a": float64(5),
			"b": float64(3),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "success", addResult.Status)
	assert.Equal(t, float64(8), addResult.Result)

	// Test subtract tool
	subtractResult, err := registry.CallTool("subtract", structures.ToolCallFunction{
		Name: "subtract",
		Arguments: map[string]interface{}{
			"a": float64(5),
			"b": float64(3),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "success", subtractResult.Status)
	assert.Equal(t, float64(2), subtractResult.Result)
}
