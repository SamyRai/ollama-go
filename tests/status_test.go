package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetVersion validates retrieving the API version.
func TestGetVersion(t *testing.T) {
	cli, rec := SetupVCRTest(t, "get_version")
	defer rec.Stop()

	resp, err := cli.GetVersion()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Version)
}

// TestGetRunningProcesses validates retrieving running model processes.
func TestGetRunningProcesses(t *testing.T) {
	cli, rec := SetupVCRTest(t, "get_running_processes")
	defer rec.Stop()

	resp, err := cli.GetRunningProcesses()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.GreaterOrEqual(t, len(resp.Models), 0)
}
