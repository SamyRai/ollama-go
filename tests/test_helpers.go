package tests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/SamyRai/ollama-go/internal/client"
	"github.com/SamyRai/ollama-go/internal/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
)

// SetupVCRTest creates a properly configured VCR recorder and client for testing.
// It ensures tests use recorded responses when available.
func SetupVCRTest(t *testing.T, fixtureName string) (*client.OllamaClient, *recorder.Recorder) {
	// Force replay mode to ensure we use recorded responses when available
	os.Setenv("VCR_MODE", "replay")

	// Get absolute path to the test directory for reliable fixture loading
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)
	fixturePath := filepath.Join(testDir, "fixtures", fixtureName)

	// Create and configure the recorder
	rec, err := recorder.New(fixturePath)
	require.NoError(t, err)

	// Log the fixture path for debugging purposes
	t.Logf("Loading VCR fixtures from: %s", fixturePath)

	// Create and configure the client to use the recorder
	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	return cli, rec
}
