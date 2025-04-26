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
type OllamaClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes a new Ollama API client with default settings.
func NewClient(cfg *config.Config) *OllamaClient {
	return &OllamaClient{
		BaseURL: cfg.BaseURL,
		HTTPClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// Request handles normal HTTP requests (non-streaming).
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
		return errors.New("API request failed with status: " + resp.Status)
	}

	// Decode the response
	return json.NewDecoder(resp.Body).Decode(response)
}

// StreamRequest handles streaming HTTP responses.
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
