package api

import "github.com/SamyRai/ollama-go/internal/structures"

// Options defines customizable parameters for model behavior.
type Options = structures.Options

// NewOptions creates a new Options instance with default values.
func NewOptions() *Options {
	return &Options{}
}

// WithTemperature sets the temperature parameter.
func WithTemperature(temperature float64) func(*Options) {
	return func(o *Options) {
		o.Temperature = temperature
	}
}

// WithTopP sets the top_p parameter.
func WithTopP(topP float64) func(*Options) {
	return func(o *Options) {
		o.TopP = topP
	}
}

// WithTopK sets the top_k parameter.
func WithTopK(topK int) func(*Options) {
	return func(o *Options) {
		o.TopK = topK
	}
}

// WithMirostat sets the mirostat parameter.
func WithMirostat(mirostat int) func(*Options) {
	return func(o *Options) {
		o.Mirostat = mirostat
	}
}

// WithMirostatTau sets the mirostat_tau parameter.
func WithMirostatTau(mirostatTau float64) func(*Options) {
	return func(o *Options) {
		o.MirostatTau = mirostatTau
	}
}

// WithMirostatEta sets the mirostat_eta parameter.
func WithMirostatEta(mirostatEta float64) func(*Options) {
	return func(o *Options) {
		o.MirostatEta = mirostatEta
	}
}

// WithRepeatPenalty sets the repeat_penalty parameter.
func WithRepeatPenalty(repeatPenalty float64) func(*Options) {
	return func(o *Options) {
		o.RepeatPenalty = repeatPenalty
	}
}

// WithRepeatLastN sets the repeat_last_n parameter.
func WithRepeatLastN(repeatLastN int) func(*Options) {
	return func(o *Options) {
		o.RepeatLastN = repeatLastN
	}
}

// WithFrequencyPenalty sets the frequency_penalty parameter.
func WithFrequencyPenalty(frequencyPenalty float64) func(*Options) {
	return func(o *Options) {
		o.FrequencyPenalty = frequencyPenalty
	}
}

// WithPresencePenalty sets the presence_penalty parameter.
func WithPresencePenalty(presencePenalty float64) func(*Options) {
	return func(o *Options) {
		o.PresencePenalty = presencePenalty
	}
}

// WithTFS sets the tfs parameter.
func WithTFS(tfs float64) func(*Options) {
	return func(o *Options) {
		o.TFS = tfs
	}
}

// WithTopA sets the top_a parameter.
func WithTopA(topA float64) func(*Options) {
	return func(o *Options) {
		o.TopA = topA
	}
}

// WithTypicalP sets the typical_p parameter.
func WithTypicalP(typicalP float64) func(*Options) {
	return func(o *Options) {
		o.TypicalP = typicalP
	}
}

// WithGrammar sets the grammar parameter.
func WithGrammar(grammar string) func(*Options) {
	return func(o *Options) {
		o.Grammar = grammar
	}
}

// ApplyOptions applies a list of option functions to an Options instance.
func ApplyOptions(o *Options, opts ...func(*Options)) {
	for _, opt := range opts {
		opt(o)
	}
}
