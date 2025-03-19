package structures

// =========================
// == Tool API ==
// =========================

// Tool defines a callable function available to models.
type Tool struct {
	Type     string       `json:"type"`     // Type of tool (e.g., "function").
	Function ToolFunction `json:"function"` // Details of the function.
}

// ToolFunction describes the function a model can invoke.
type ToolFunction struct {
	Name        string               `json:"name"`        // Function name.
	Description string               `json:"description"` // Function description.
	Parameters  map[string]ToolParam `json:"parameters"`  // Function parameters.
}

// ToolParam defines a function parameter.
type ToolParam struct {
	Type        string   `json:"type"`           // Parameter data type.
	Description string   `json:"description"`    // Parameter description.
	Enum        []string `json:"enum,omitempty"` // Optional: Allowed values.
}

// ToolCall represents an invocation of a tool.
type ToolCall struct {
	Function ToolCallFunction `json:"function"` // Function call details.
}

// ToolCallFunction contains function call arguments.
type ToolCallFunction struct {
	Name      string                 `json:"name"`      // Function name.
	Arguments map[string]interface{} `json:"arguments"` // Function arguments.
}

// ToolCallResult represents the result of executing a tool.
type ToolCallResult struct {
	Status   string                 `json:"status"`             // The status of the tool call ("success", "failure", etc.).
	Result   interface{}            `json:"result"`             // The result of the tool call, could be any type based on the tool.
	Error    string                 `json:"error,omitempty"`    // Optional: Error message if tool execution fails.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Optional: Metadata about the result.
}
