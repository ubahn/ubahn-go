package v2

type flowOutput struct {
	ExpectedInputs map[string]string `yaml:"expectedInputs"`
	Fallback       string            `yaml:"fallback"`
	empty          bool
}

type flowConfig struct {
	Sequence []string              `yaml:"sequence"`
	Triggers []string              `yaml:"triggers"`
	Fallback string                `yaml:"fallback"`
	Outputs  map[string]flowOutput `yaml:"outputs"`
}
