// Package config provides configuration management for the Ollama Go client.
package config

import (
	"os"
	"strings"
	"time"
)

// Config holds the client configuration settings.
type Config struct {
	BaseURL string        // API Base URL
	Timeout time.Duration // Request timeout duration
	APIKey  string        // Authentication key (if required in the future)
	Debug   bool          // Enable debug logging
}

// DefaultConfig returns a default configuration.
func DefaultConfig() *Config {
	// Determine if debug mode should be enabled from environment
	debugMode := false
	debugEnv := strings.ToLower(os.Getenv("OLLAMA_DEBUG"))
	if debugEnv == "true" || debugEnv == "1" || debugEnv == "yes" {
		debugMode = true
	}

	return &Config{
		BaseURL: "http://localhost:11434",
		Timeout: 30 * time.Second,            // 30-second timeout for API requests
		APIKey:  os.Getenv("OLLAMA_API_KEY"), // Load API Key from environment variable (if needed)
		Debug:   debugMode,                   // Enable debug logging based on environment
	}
}
