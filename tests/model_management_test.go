package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/SamyRai/ollama-go/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateModel validates model creation.
func TestCreateModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var reqBody structures.ModelManagementRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		assert.Equal(t, "test-model", reqBody.Name)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	req := structures.ModelManagementRequest{Name: "test-model"}
	err := cli.CreateModel(req)
	require.NoError(t, err)
}

// TestDeleteModel validates model deletion.
func TestDeleteModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		var reqBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		assert.Equal(t, "test-model", reqBody["name"])
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.DeleteModel("test-model")
	require.NoError(t, err)
}

// TestDeleteModelFailure validates model deletion failure.
func TestDeleteModelFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.DeleteModel("test-model")
	require.Error(t, err)
	require.Contains(t, err.Error(), "404")
}

// TestCopyModel validates model copying.
func TestCopyModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var reqBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		assert.Equal(t, "test-model", reqBody["source"])
		assert.Equal(t, "test-model-copy", reqBody["destination"])
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.CopyModel("test-model", "test-model-copy")
	require.NoError(t, err)
}

// TestCopyModelFailure validates model copying failure.
func TestCopyModelFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.CopyModel("test-model", "test-model-copy")
	require.Error(t, err)
	require.Contains(t, err.Error(), "400")
}

// TestPullModel validates pulling a model from a remote source.
func TestPullModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var reqBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		assert.Equal(t, "test-model", reqBody["name"])
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.PullModel("test-model")
	require.NoError(t, err)
}

// TestPushModel validates pushing a model to a remote source.
func TestPushModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var reqBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		assert.Equal(t, "test-model", reqBody["name"])
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cli := client.NewClient(&config.Config{BaseURL: server.URL})
	err := cli.PushModel("test-model")
	require.NoError(t, err)
}
