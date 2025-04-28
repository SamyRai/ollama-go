package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/SamyRai/ollama-go/internal/utils"
)

// OllamaClient provides a structured API client for communicating with the Ollama API.
type OllamaClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Logger     *utils.Logger
}

// NewClient initializes a new Ollama API client with default settings.
func NewClient(cfg *config.Config) *OllamaClient {
	logger := utils.GetLogger()
	if cfg.Debug {
		logger.SetLevel(utils.LogLevelDebug)
	}

	return &OllamaClient{
		BaseURL: cfg.BaseURL,
		HTTPClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		Logger: logger,
	}
}

// Request handles normal HTTP requests (non-streaming).
func (c *OllamaClient) Request(method, endpoint string, body interface{}, response interface{}) error {
	url := c.BaseURL + endpoint
	c.Logger.Debug("Making %s request to %s", method, url)

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			c.Logger.Error("Failed to marshal request body: %v", err)
			return err
		}

		if c.Logger.DebugEnabled() {
			c.Logger.Debug("Request body: %s", string(reqBody))
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		c.Logger.Error("Failed to create request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	c.Logger.Debug("Sending request to %s", url)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Error("Request failed: %v", err)
		return err
	}
	defer resp.Body.Close()

	c.Logger.Debug("Received response with status: %s", resp.Status)

	if resp.StatusCode >= 400 {
		errMsg := fmt.Sprintf("API request failed with status: %s", resp.Status)
		c.Logger.Error("%s", errMsg)
		return errors.New(errMsg)
	}

	// Decode the response if a response object was provided
	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			c.Logger.Error("Failed to decode response: %v", err)
			return err
		}
		c.Logger.Debug("Successfully decoded response")
	} else {
		// If no response object was provided, just read and discard the body
		if _, err = io.Copy(io.Discard, resp.Body); err != nil {
			c.Logger.Error("Failed to read response body: %v", err)
			return err
		}
	}

	return nil
}

// StreamRequest handles streaming HTTP responses.
func (c *OllamaClient) StreamRequest(method, endpoint string, body interface{}, callback func(json.RawMessage)) error {
	url := c.BaseURL + endpoint
	c.Logger.Debug("Making streaming %s request to %s", method, url)

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			c.Logger.Error("Failed to marshal request body: %v", err)
			return err
		}

		if c.Logger.DebugEnabled() {
			c.Logger.Debug("Request body: %s", string(reqBody))
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		c.Logger.Error("Failed to create request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	c.Logger.Debug("Sending streaming request to %s", url)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Error("Streaming request failed: %v", err)
		return err
	}
	defer resp.Body.Close()

	c.Logger.Debug("Received streaming response with status: %s", resp.Status)

	if resp.StatusCode >= 400 {
		errMsg := fmt.Sprintf("API streaming request failed with status: %s", resp.Status)
		c.Logger.Error("%s", errMsg)
		return errors.New(errMsg)
	}

	// Process the streaming response
	c.Logger.Debug("Starting to process stream")
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				c.Logger.Debug("End of stream reached")
				break // End of stream
			}
			c.Logger.Error("Error reading stream: %v", err)
			return err
		}

		// Process JSON chunk
		var message json.RawMessage
		if err := json.Unmarshal(line, &message); err != nil {
			c.Logger.Error("Failed to unmarshal stream message: %v", err)
			return err
		}

		if c.Logger.DebugEnabled() {
			c.Logger.Debug("Stream chunk: %s", string(line))
		}

		callback(message)
	}

	c.Logger.Debug("Streaming request completed successfully")
	return nil
}
