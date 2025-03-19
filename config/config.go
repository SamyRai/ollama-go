package config

import (
    "os"
    "time"
)

// Config holds the client configuration settings.
type Config struct {
    BaseURL    string        // API Base URL
    Timeout    time.Duration // Request timeout duration
    APIKey     string        // Authentication key (if required in the future)
}

// DefaultConfig returns a default configuration.
func DefaultConfig() *Config {
    return &Config{
        BaseURL: "http://localhost:11434",
        Timeout: 30 * time.Second, // 30-second timeout for API requests
        APIKey:  os.Getenv("OLLAMA_API_KEY"), // Load API Key from environment variable (if needed)
    }
}
