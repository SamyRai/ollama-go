package tests

import (
	"testing"

	"github.com/SamyRai/ollama-go/api"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	// Test creating new options
	options := api.NewOptions()
	assert.NotNil(t, options, "Options should not be nil")

	// Test default values
	assert.Zero(t, options.Temperature, "Default temperature should be zero")
	assert.Zero(t, options.TopP, "Default top_p should be zero")
	assert.Zero(t, options.TopK, "Default top_k should be zero")

	// Test individual option functions
	tests := []struct {
		name     string
		option   func(*api.Options)
		validate func(*testing.T, *api.Options)
	}{
		{
			name:   "WithTemperature",
			option: api.WithTemperature(0.7),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.7, o.Temperature, "Temperature should be set to 0.7")
			},
		},
		{
			name:   "WithTopP",
			option: api.WithTopP(0.9),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.9, o.TopP, "TopP should be set to 0.9")
			},
		},
		{
			name:   "WithTopK",
			option: api.WithTopK(40),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 40, o.TopK, "TopK should be set to 40")
			},
		},
		{
			name:   "WithMirostat",
			option: api.WithMirostat(2),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 2, o.Mirostat, "Mirostat should be set to 2")
			},
		},
		{
			name:   "WithMirostatTau",
			option: api.WithMirostatTau(5.0),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 5.0, o.MirostatTau, "MirostatTau should be set to 5.0")
			},
		},
		{
			name:   "WithMirostatEta",
			option: api.WithMirostatEta(0.1),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.1, o.MirostatEta, "MirostatEta should be set to 0.1")
			},
		},
		{
			name:   "WithRepeatPenalty",
			option: api.WithRepeatPenalty(1.1),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 1.1, o.RepeatPenalty, "RepeatPenalty should be set to 1.1")
			},
		},
		{
			name:   "WithRepeatLastN",
			option: api.WithRepeatLastN(64),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 64, o.RepeatLastN, "RepeatLastN should be set to 64")
			},
		},
		{
			name:   "WithFrequencyPenalty",
			option: api.WithFrequencyPenalty(0.5),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.5, o.FrequencyPenalty, "FrequencyPenalty should be set to 0.5")
			},
		},
		{
			name:   "WithPresencePenalty",
			option: api.WithPresencePenalty(0.5),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.5, o.PresencePenalty, "PresencePenalty should be set to 0.5")
			},
		},
		{
			name:   "WithTFS",
			option: api.WithTFS(0.5),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.5, o.TFS, "TFS should be set to 0.5")
			},
		},
		{
			name:   "WithTopA",
			option: api.WithTopA(0.5),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.5, o.TopA, "TopA should be set to 0.5")
			},
		},
		{
			name:   "WithTypicalP",
			option: api.WithTypicalP(0.5),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, 0.5, o.TypicalP, "TypicalP should be set to 0.5")
			},
		},
		{
			name:   "WithGrammar",
			option: api.WithGrammar("json"),
			validate: func(t *testing.T, o *api.Options) {
				assert.Equal(t, "json", o.Grammar, "Grammar should be set to 'json'")
			},
		},
	}

	// Test each option individually
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := api.NewOptions()
			tt.option(o)
			tt.validate(t, o)
		})
	}

	// Test applying multiple options
	t.Run("ApplyOptions", func(t *testing.T) {
		o := api.NewOptions()
		api.ApplyOptions(o,
			api.WithTemperature(0.7),
			api.WithTopP(0.9),
			api.WithTopK(40),
			api.WithRepeatPenalty(1.1),
		)

		assert.Equal(t, 0.7, o.Temperature, "Temperature should be set to 0.7")
		assert.Equal(t, 0.9, o.TopP, "TopP should be set to 0.9")
		assert.Equal(t, 40, o.TopK, "TopK should be set to 40")
		assert.Equal(t, 1.1, o.RepeatPenalty, "RepeatPenalty should be set to 1.1")
	})
}
