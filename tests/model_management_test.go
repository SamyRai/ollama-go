package tests

import (
	"github.com/SamyRai/ollama-go/client"
	"github.com/SamyRai/ollama-go/config"
	"github.com/SamyRai/ollama-go/structures"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
	"testing"
)

// TestCreateModel validates model creation.
func TestCreateModel(t *testing.T) {
	rec, err := recorder.New("fixtures/create_model")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	req := structures.ModelManagementRequest{Name: "test-model"}

	err = cli.CreateModel(req)

	require.NoError(t, err)
}

// TestDeleteModel validates model deletion.
func TestDeleteModel(t *testing.T) {
	rec, err := recorder.New("fixtures/delete_model")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	err = cli.DeleteModel("test-model")

	require.NoError(t, err)
}

// TestCopyModel validates model copying.
func TestCopyModel(t *testing.T) {
	rec, err := recorder.New("fixtures/copy_model")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	err = cli.CopyModel("test-model", "test-model-copy")

	require.NoError(t, err)
}

// TestPullModel validates pulling a model from a remote source.
func TestPullModel(t *testing.T) {
	rec, err := recorder.New("fixtures/pull_model")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	err = cli.PullModel("test-model")

	require.NoError(t, err)
}

// TestPushModel validates pushing a model to a remote source.
func TestPushModel(t *testing.T) {
	rec, err := recorder.New("fixtures/push_model")
	require.NoError(t, err)
	defer rec.Stop()

	cli := client.NewClient(config.DefaultConfig())
	cli.HTTPClient.Transport = rec

	err = cli.PushModel("test-model")

	require.NoError(t, err)
}
