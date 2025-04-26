package testutils

import (
	"path/filepath"
	"testing"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
)

// SetupVCR creates a new VCR recorder and configures a client to use it.
// It returns the client and a cleanup function that should be deferred.
func SetupVCR(t *testing.T, cassetteName string) (*client.OllamaClient, func()) {
	t.Helper()

	// Normalize the cassette path
	fixturePath := filepath.Join("testdata", "fixtures", cassetteName)

	// Create a new recorder
	rec, err := recorder.New(fixturePath)
	require.NoError(t, err, "Failed to create VCR recorder")

	// Create a new client with the recorder as transport
	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	// Return the client and a cleanup function
	return cli, func() {
		err := rec.Stop()
		require.NoError(t, err, "Failed to stop VCR recorder")
	}
}

// SetupVCRWithMode creates a new VCR recorder with a specific mode and configures a client to use it.
// Mode can be recorder.ModeRecording or recorder.ModeReplaying.
func SetupVCRWithMode(t *testing.T, cassetteName string, mode recorder.Mode) (*client.OllamaClient, func()) {
	t.Helper()

	// Normalize the cassette path
	fixturePath := filepath.Join("testdata", "fixtures", cassetteName)

	// Create a new recorder with the specified mode
	rec, err := recorder.NewAsMode(fixturePath, mode, nil)
	require.NoError(t, err, "Failed to create VCR recorder")

	// Create a new client with the recorder as transport
	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	// Return the client and a cleanup function
	return cli, func() {
		err := rec.Stop()
		require.NoError(t, err, "Failed to stop VCR recorder")
	}
}
