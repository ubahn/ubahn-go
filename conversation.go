package ubahn

// IConversation defines a conversation that can be continued.
type IConversation interface {
	// Continue finds the next output to the given previous output and input.
	Continue(prevOutput IOutput, input IInput) string
}

// Conversation is the default implementation of IConversation.
// It's continued based on the sequence, defined in the give YAML file.
type Conversation struct {
	Empty    bool
	v1Config *conversationConfigV1
	file     *conversationFile
}

// NewConversation creates a new instance of a conversation,
// initialized from the given YAML file.
// If initialization filed, an error is returned.
func NewConversation(conversationFilePath string) (*Conversation, error) {
	nullConv := &Conversation{Empty: true}
	file, err := newConversationFile(conversationFilePath)
	if err != nil {
		return nullConv, err
	}
	conv := &Conversation{file: file}
	if file.version == 1 {
		conv.v1Config, err = file.V1Config()
		if err != nil {
			return nullConv, err
		}
	}
	return conv, nil
}

// Continue finds the next output to the given previous output and input.
// If no suitable output found, it returns blank output.
func (conv *Conversation) Continue(prevOutput IOutput, input IInput) string {
	if prevOutput.Name() == BlankOutputName && conv.inTriggers(input) {
		return conv.firstOutputName()
	}
	return conv.matchOutput(prevOutput, input)
}

func (conv *Conversation) inTriggers(input IInput) bool {
	if conv.useV1() {
		return conv.inTriggersV1(input)
	}
	return false
}

func (conv *Conversation) firstOutputName() string {
	if conv.useV1() {
		return conv.firstOutputNameV1()
	}
	return BlankOutputName
}

func (conv *Conversation) matchOutput(prevOutput IOutput, input IInput) string {
	if conv.useV1() {
		return conv.matchOutputV1(prevOutput, input)
	}
	return BlankOutputName
}

func (conv *Conversation) useV1() bool {
	return conv.file.version == 1 && conv.v1Config != nil
}

func (conv *Conversation) inTriggersV1(input IInput) bool {
	triggers := conv.v1Config.Triggers
	for i := 0; i < len(triggers); i++ {
		trigger := triggers[i]
		if input.Name() == trigger {
			return true
		}
	}
	return false
}

func (conv *Conversation) firstOutputNameV1() string {
	sequence := conv.v1Config.Sequence
	if len(sequence) > 0 {
		return sequence[0]
	}
	return conv.fallbackV1()
}

func (conv *Conversation) fallbackV1() string {
	fallback := conv.v1Config.Fallback
	if len(fallback) > 0 {
		return fallback
	}
	return BlankOutputName
}

func (conv *Conversation) matchOutputV1(prevOutput IOutput, input IInput) string {
	prevOutputConfig := conv.findPrevOutputConfigV1(prevOutput)
	if prevOutputConfig.empty {
		return conv.fallbackV1()
	}
	output := prevOutputConfig.ExpectedInputs[input.Name()]
	if len(output) == 0 {
		if len(prevOutputConfig.Fallback) > 0 {
			return prevOutputConfig.Fallback
		}
		return conv.fallbackV1()
	}
	if output == NextOutputName {
		return conv.findNextOutputV1(prevOutput.Name())
	}
	return output
}

func (conv *Conversation) findPrevOutputConfigV1(prevOutput IOutput) configV1Output {
	sequence := conv.v1Config.Sequence
	for i := 0; i < len(sequence); i++ {
		outputName := sequence[i]
		if prevOutput.Name() == outputName {
			return conv.v1Config.Outputs[outputName]
		}
	}
	return configV1Output{empty: true}
}

func (conv *Conversation) findNextOutputV1(prevOutputName string) string {
	sequence := conv.v1Config.Sequence
	for i := 0; i < len(sequence)-1; i++ {
		output := sequence[i]
		if prevOutputName == output {
			return sequence[i+1]
		}
	}
	return conv.fallbackV1()
}
