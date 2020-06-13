package ubahn

type flowOutput struct {
	ExpectedInputs map[string]string `yaml:"expectedInputs"`
	Fallback       string            `yaml:"fallback"`
	Exit           bool              `yaml:"exit"`
	empty          bool
}

type flowConfig struct {
	Fallback   string                  `yaml:"fallback"`
	RootOutput string                  `yaml:"rootOutput"`
	Outputs    []map[string]flowOutput `yaml:"outputs"`
}
