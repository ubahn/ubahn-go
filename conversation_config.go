package ubahn

type conversationConfig struct {
	Triggers       map[string]string `yaml:"triggers"`
	DefaultTrigger string            `yaml:"defaultTrigger"`
}
