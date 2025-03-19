package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
	"hrelay/core/llm/ollama/client"
	"hrelay/core/llm/ollama/config"
	"testing"
)

// TestGetVersion validates retrieving the API version.
func TestGetVersion(t *testing.T) {
	rec, err := recorder.New("fixtures/get_version")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	resp, err := cli.GetVersion()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Version)
}

// TestGetRunningProcesses validates retrieving running model processes.
func TestGetRunningProcesses(t *testing.T) {
	rec, err := recorder.New("fixtures/get_running_processes")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	resp, err := cli.GetRunningProcesses()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.GreaterOrEqual(t, len(resp.Models), 0)
}
