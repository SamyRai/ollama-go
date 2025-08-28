package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/SamyRai/ollama-go/config"
	"io"
	"net/http"
)

// OllamaClient provides a structured API client for communicating with the Ollama API.
// It handles the construction of HTTP requests and the parsing of responses.
type OllamaClient struct {
	BaseURL    string       // The base URL of the Ollama API (e.g., "http://localhost:11434").
	HTTPClient *http.Client // The HTTP client to use for making requests.
}

// NewClient initializes a new Ollama API client with the provided configuration.
// It returns a pointer to a new OllamaClient.
func NewClient(cfg *config.Config) *OllamaClient {
	return &OllamaClient{
		BaseURL: cfg.BaseURL,
		HTTPClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// Request sends a standard, non-streaming HTTP request to the Ollama API.
// It marshals the request body to JSON, sends the request, and decodes the JSON response.
// If the response has a status code of 400 or greater, it returns an error.
func (c *OllamaClient) Request(method, endpoint string, body interface{}, response interface{}) error {
	url := c.BaseURL + endpoint

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return errors.New("API request failed with status: " + resp.Status + " and body: " + string(bodyBytes))
	}

	// Decode the response if a response container is provided
	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}

	return nil
}

// StreamRequest sends a streaming HTTP request to the Ollama API.
// It processes the response as a stream of JSON objects, calling the provided
// callback function for each object.
func (c *OllamaClient) StreamRequest(method, endpoint string, body interface{}, callback func(json.RawMessage)) error {
	url := c.BaseURL + endpoint

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("API request failed with status: " + resp.Status)
	}

	// Process the streaming response
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break // End of stream
			}
			return err
		}

		// Process JSON chunk
		var message json.RawMessage
		if err := json.Unmarshal(line, &message); err != nil {
			return err
		}

		callback(message)
	}

	return nil
}
