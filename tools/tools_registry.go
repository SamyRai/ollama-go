package tools

import (
	"errors"
	"hrelay/core/llm/ollama/structures"
	"sync"
)

// ToolRegistry manages registered tools with strict function definitions.
type ToolRegistry struct {
	mu    sync.RWMutex
	tools map[string]func(args structures.ToolCallFunction) (structures.ToolCallResult, error)
}

// NewRegistry initializes a strict tool registry.
func NewRegistry() *ToolRegistry {
	return &ToolRegistry{
		tools: make(map[string]func(args structures.ToolCallFunction) (structures.ToolCallResult, error)),
	}
}

// RegisterTool registers a tool function by strict definition.
func (r *ToolRegistry) RegisterTool(name string, handler func(args structures.ToolCallFunction) (structures.ToolCallResult, error)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tools[name] = handler
}

// CallTool executes a registered tool function.
func (r *ToolRegistry) CallTool(name string, args structures.ToolCallFunction) (structures.ToolCallResult, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if fn, exists := r.tools[name]; exists {
		return fn(args)
	}
	return structures.ToolCallResult{}, errors.New("tool not registered: " + name)
}
