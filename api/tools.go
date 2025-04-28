package api

import (
	"github.com/SamyRai/ollama-go/internal/structures"
	"github.com/SamyRai/ollama-go/internal/tools"
)

// ToolCallResult represents the result of executing a tool.
type ToolCallResult = structures.ToolCallResult

// ToolRegistry manages registered tools with strict function definitions.
type ToolRegistry struct {
	registry *tools.ToolRegistry
}

// NewToolRegistry initializes a new tool registry.
func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{
		registry: tools.NewRegistry(),
	}
}

// RegisterTool registers a tool function by name.
func (r *ToolRegistry) RegisterTool(name string, handler func(args map[string]interface{}) (interface{}, error)) {
	r.registry.RegisterTool(name, func(args structures.ToolCallFunction) (structures.ToolCallResult, error) {
		result, err := handler(args.Arguments)
		if err != nil {
			return structures.ToolCallResult{
				Status: "error",
				Error:  err.Error(),
			}, nil
		}
		return structures.ToolCallResult{
			Status: "success",
			Result: result,
		}, nil
	})
}

// CallTool executes a registered tool function.
func (r *ToolRegistry) CallTool(name string, args map[string]interface{}) (interface{}, error) {
	toolArgs := structures.ToolCallFunction{
		Name:      name,
		Arguments: args,
	}
	result, err := r.registry.CallTool(name, toolArgs)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}
