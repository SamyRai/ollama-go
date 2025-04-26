package testutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/SamyRai/ollama-go/internal/structures"
)

// MockServer provides a test server that simulates Ollama API responses.
type MockServer struct {
	Server *httptest.Server
}

// NewMockServer creates a new mock Ollama API server.
func NewMockServer() *MockServer {
	mux := http.NewServeMux()

	// Setup routes
	setupChatRoute(mux)
	setupCompletionRoute(mux)
	setupEmbeddingsRoute(mux)
	setupModelRoutes(mux)
	setupStatusRoutes(mux)

	// Create the server
	server := httptest.NewServer(mux)

	return &MockServer{
		Server: server,
	}
}

// Close shuts down the mock server.
func (m *MockServer) Close() {
	m.Server.Close()
}

// URL returns the base URL of the mock server.
func (m *MockServer) URL() string {
	return m.Server.URL
}

// Setup chat route
func setupChatRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/chat", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req structures.ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create a response
		resp := structures.ChatResponse{
			Model:     req.Model,
			CreatedAt: time.Now(),
			Message: structures.Message{
				Role:    "assistant",
				Content: "This is a mock response from the test server.",
			},
			Done: true,
		}

		// If the request includes tools, add a tool call to the response
		if len(req.Tools) > 0 {
			for _, tool := range req.Tools {
				if tool.Type == "function" && tool.Function.Name == "getWeather" {
					resp.Message.ToolCalls = []structures.ToolCall{
						{
							Function: structures.ToolCallFunction{
								Name: "getWeather",
								Arguments: map[string]interface{}{
									"location": "Paris",
								},
							},
						},
					}
					break
				}
			}
		}

		// If streaming is requested, send multiple responses
		if req.Stream {
			// Set the appropriate headers for streaming
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			// Send multiple chunks
			for i := 0; i < 3; i++ {
				chunkResp := resp
				chunkResp.Message.Content = fmt.Sprintf("Chunk %d of the response. ", i+1)
				chunkResp.Done = (i == 2) // Only the last chunk is done

				// Encode and send the chunk
				if err := json.NewEncoder(w).Encode(chunkResp); err != nil {
					http.Error(w, "Error encoding response", http.StatusInternalServerError)
					return
				}
				w.(http.Flusher).Flush()
				time.Sleep(10 * time.Millisecond) // Small delay between chunks
			}
			return
		}

		// For non-streaming, just send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})
}

// Setup completion route
func setupCompletionRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/generate", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req structures.CompletionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create a response
		resp := structures.CompletionResponse{
			Model:     req.Model,
			CreatedAt: time.Now(),
			Response:  "This is a mock completion response from the test server.",
			Done:      true,
		}

		// If streaming is requested, send multiple responses
		if req.Stream {
			// Set the appropriate headers for streaming
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			// Send multiple chunks
			for i := 0; i < 3; i++ {
				chunkResp := resp
				chunkResp.Response = fmt.Sprintf("Chunk %d of the completion. ", i+1)
				chunkResp.Done = (i == 2) // Only the last chunk is done

				// Encode and send the chunk
				if err := json.NewEncoder(w).Encode(chunkResp); err != nil {
					http.Error(w, "Error encoding response", http.StatusInternalServerError)
					return
				}
				w.(http.Flusher).Flush()
				time.Sleep(10 * time.Millisecond) // Small delay between chunks
			}
			return
		}

		// For non-streaming, just send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})
}

// Setup embeddings route
func setupEmbeddingsRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/embed", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req structures.EmbeddingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create a response with mock embeddings
		resp := structures.EmbeddingResponse{
			Model: req.Model,
		}

		// Generate mock embeddings for each input
		for range req.Input {
			// Create a mock embedding vector with 10 dimensions
			embedding := make([]float32, 10)
			for i := range embedding {
				embedding[i] = float32(i) * 0.1 // Simple pattern for mock data
			}
			resp.Embeddings = append(resp.Embeddings, embedding)
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})
}

// Setup model routes
func setupModelRoutes(mux *http.ServeMux) {
	// List models
	mux.HandleFunc("/api/tags", func(w http.ResponseWriter, r *http.Request) {
		// Only accept GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create a response with mock models
		resp := structures.ModelListResponse{
			Models: []structures.ModelInfo{
				{
					Name:       "llama3",
					Size:       10000000000, // 10GB
					ModifiedAt: time.Now(),
					Digest:     "sha256:abc123",
					Details:    structures.ModelDetails{},
				},
				{
					Name:       "mistral",
					Size:       8000000000, // 8GB
					ModifiedAt: time.Now(),
					Digest:     "sha256:def456",
					Details:    structures.ModelDetails{},
				},
			},
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Show model
	mux.HandleFunc("/api/show", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req structures.ShowModelRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create a response
		resp := structures.ShowModelResponse{
			Name:        req.Model,
			Version:     "v1.0.0",
			Description: "This is a mock model description.",
			Tags:        []string{"mock", "test"},
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Create model
	mux.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req structures.ModelManagementRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Send an empty response for success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// Delete model
	mux.HandleFunc("/api/delete", func(w http.ResponseWriter, r *http.Request) {
		// Only accept DELETE requests
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if model name is provided
		if _, ok := req["model"]; !ok {
			http.Error(w, "Model name is required", http.StatusBadRequest)
			return
		}

		// Send an empty response for success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// Copy model
	mux.HandleFunc("/api/copy", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if source and destination are provided
		if _, ok := req["sourceModel"]; !ok {
			http.Error(w, "Source model name is required", http.StatusBadRequest)
			return
		}
		if _, ok := req["targetModel"]; !ok {
			http.Error(w, "Target model name is required", http.StatusBadRequest)
			return
		}

		// Send an empty response for success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// Pull model
	mux.HandleFunc("/api/pull", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if model name is provided
		if _, ok := req["model"]; !ok {
			http.Error(w, "Model name is required", http.StatusBadRequest)
			return
		}

		// Send an empty response for success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// Push model
	mux.HandleFunc("/api/push", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if model name is provided
		if _, ok := req["model"]; !ok {
			http.Error(w, "Model name is required", http.StatusBadRequest)
			return
		}

		// Send an empty response for success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

// Setup status routes
func setupStatusRoutes(mux *http.ServeMux) {
	// Version endpoint
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		// Only accept GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create a response
		resp := structures.VersionResponse{
			Version: "0.1.0",
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Running processes endpoint
	mux.HandleFunc("/api/ps", func(w http.ResponseWriter, r *http.Request) {
		// Only accept GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create a response with mock running processes
		resp := structures.ModelProcessResponse{
			Models: []structures.ModelProcess{
				{
					Name:      "llama3",
					Model:     "llama3",
					Size:      10000000000, // 10GB
					VRAMSize:  5000000000,  // 5GB
					Digest:    "sha256:abc123",
					ExpiresAt: time.Now().Add(24 * time.Hour),
					Details:   structures.ModelDetails{},
				},
			},
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})
}
