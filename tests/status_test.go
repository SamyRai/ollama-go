package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetVersion validates retrieving the API version.
func TestGetVersion(t *testing.T) {
	t.Skip("Skipping test that requires a running Ollama server to re-record cassettes.")
	cli, _ := SetupVCRTest(t, "get_version")

	resp, err := cli.GetVersion()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Version)
}

// TestGetRunningProcesses validates retrieving running model processes.
func TestGetRunningProcesses(t *testing.T) {
	t.Skip("Skipping test that requires a running Ollama server to re-record cassettes.")
	cli, _ := SetupVCRTest(t, "get_running_processes")

	resp, err := cli.GetRunningProcesses()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.GreaterOrEqual(t, len(resp.Models), 0)
}
