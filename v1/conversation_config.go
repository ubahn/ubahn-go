package v1

type configOutput struct {
	ExpectedInputs map[string]string `yaml:"expectedInputs"`
	Fallback       string            `yaml:"fallback"`
	empty          bool
}

type conversationConfig struct {
	Sequence []string                `yaml:"sequence"`
	Triggers []string                `yaml:"triggers"`
	Fallback string                  `yaml:"fallback"`
	Outputs  map[string]configOutput `yaml:"outputs"`
}
