package v1

import (
	core "github.com/ubahn/ubahn-go/core"
)

// Conversation is the default implementation of IConversation.
// It's continued based on the sequence, defined in the give YAML file.
type Conversation struct {
	config conversationConfig
}

// NewConversation creates a new instance of a conversation.
func NewConversation(file core.IConversationFile) (core.IConversation, error) {
	conv := &Conversation{}
	err := file.Parse(&conv.config)
	if err != nil {
		return core.NullConversation, err
	}
	return conv, nil
}

// Continue finds the next output to the given previous output and input.
// If no suitable output found, it returns blank output.
func (conv *Conversation) Continue(prevOutput core.IOutput, input core.IInput) string {
	if prevOutput.Name() == core.BlankOutputName && conv.inTriggers(input) {
		return conv.firstOutputName()
	}
	return conv.matchOutput(prevOutput, input)
}

// Empty returnes
func (conv *Conversation) Empty() bool {
	return false
}

func (conv *Conversation) inTriggers(input core.IInput) bool {
	triggers := conv.config.Triggers
	for i := 0; i < len(triggers); i++ {
		trigger := triggers[i]
		if input.Name() == trigger {
			return true
		}
	}
	return false
}

func (conv *Conversation) firstOutputName() string {
	sequence := conv.config.Sequence
	if len(sequence) > 0 {
		return sequence[0]
	}
	return conv.fallback()
}

func (conv *Conversation) matchOutput(prevOutput core.IOutput, input core.IInput) string {
	prevOutputConfig := conv.findPrevOutputConfig(prevOutput)
	if prevOutputConfig.empty {
		return conv.fallback()
	}
	output := prevOutputConfig.ExpectedInputs[input.Name()]
	if len(output) == 0 {
		if len(prevOutputConfig.Fallback) > 0 {
			return prevOutputConfig.Fallback
		}
		return conv.fallback()
	}
	if output == core.NextOutputName {
		return conv.findNextOutput(prevOutput.Name())
	}
	return output
}

func (conv *Conversation) fallback() string {
	fallback := conv.config.Fallback
	if len(fallback) > 0 {
		return fallback
	}
	return core.BlankOutputName
}

func (conv *Conversation) findPrevOutputConfig(prevOutput core.IOutput) configOutput {
	sequence := conv.config.Sequence
	for i := 0; i < len(sequence); i++ {
		outputName := sequence[i]
		if prevOutput.Name() == outputName {
			return conv.config.Outputs[outputName]
		}
	}
	return configOutput{empty: true}
}

func (conv *Conversation) findNextOutput(prevOutputName string) string {
	sequence := conv.config.Sequence
	for i := 0; i < len(sequence)-1; i++ {
		output := sequence[i]
		if prevOutputName == output {
			return sequence[i+1]
		}
	}
	return conv.fallback()
}
