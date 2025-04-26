package tests

import (
	"errors"
	"github.com/SamyRai/ollama-go/structures"
	"github.com/SamyRai/ollama-go/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToolRegistry(t *testing.T) {
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
				result := args.Arguments["a"].(int) * args.Arguments["b"].(int)
				return structures.ToolCallResult{
					Status: "success",
					Result: result,
				}, nil
			},
			args:      structures.ToolCallFunction{Arguments: map[string]interface{}{"a": 3, "b": 4}},
			expectErr: false,
			expected:  structures.ToolCallResult{Status: "success", Result: 12},
		},
		{
			name:      "Tool not registered should return error",
			toolName:  "nonexistent",
			toolFunc:  nil,
			args:      structures.ToolCallFunction{},
			expectErr: true,
		},
		{
			name:     "Test error handling in tool",
			toolName: "divide",
			toolFunc: func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
				if args.Arguments["b"].(int) == 0 {
					return structures.ToolCallResult{}, errors.New("division by zero")
				}
				result := args.Arguments["a"].(int) / args.Arguments["b"].(int)
				return structures.ToolCallResult{Status: "success", Result: result}, nil
			},
			args:      structures.ToolCallFunction{Arguments: map[string]interface{}{"a": 5, "b": 0}},
			expectErr: true,
		},
	}

	// Initialize ToolRegistry
	registry := tools.NewRegistry()

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.toolFunc != nil {
				// Register tool before calling
				registry.RegisterTool(tt.toolName, tt.toolFunc)
			}

			// Call the tool
			result, err := registry.CallTool(tt.toolName, tt.args)

			// Validate the result
			if tt.expectErr {
				require.Error(t, err)
				assert.Equal(t, structures.ToolCallResult{}, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
