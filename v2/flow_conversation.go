package v2

import (
	core "github.com/ubahn/ubahn-go/core"
)

// FlowConversation is the implementation of IConversation, specific to a flow.
// It's continued based on the sequence, defined in the give YAML file.
type FlowConversation struct {
	config        flowConfig
	outputFactory core.IOutputFactory
}

// NewFlowConversation creates a new instance of a flow conversation.
func NewFlowConversation(file core.IConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	conv := &FlowConversation{outputFactory: outputFactory}
	err := file.Parse(&conv.config)
	if err != nil {
		return core.NullConversation, err
	}
	return conv, nil
}

// Continue finds the next output to the given previous output and input.
// If no suitable output found, it returns blank output.
func (conv *FlowConversation) Continue(prevOutput core.IOutput, input core.IInput) core.IOutput {
	var nextOutputName string
	if prevOutput.Name() == core.BlankOutputName && conv.inTriggers(input) {
		nextOutputName = conv.firstOutputName()
	} else {
		nextOutputName = conv.matchOutput(prevOutput, input)
	}
	return conv.outputFactory.Create(nextOutputName)
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *FlowConversation) Empty() bool {
	return false
}

func (conv *FlowConversation) inTriggers(input core.IInput) bool {
	triggers := conv.config.Triggers
	for i := 0; i < len(triggers); i++ {
		trigger := triggers[i]
		if input.Name() == trigger {
			return true
		}
	}
	return false
}

func (conv *FlowConversation) firstOutputName() string {
	sequence := conv.config.Sequence
	if len(sequence) > 0 {
		return sequence[0]
	}
	return conv.fallback()
}

func (conv *FlowConversation) matchOutput(prevOutput core.IOutput, input core.IInput) string {
	prevOutputConfig := conv.findPrevOutputConfig(prevOutput)
	if prevOutputConfig.empty {
		return conv.resolveOutputName(conv.fallback(), prevOutput.Name())
	}
	output := prevOutputConfig.ExpectedInputs[input.Name()]
	if len(output) == 0 {
		if len(prevOutputConfig.Fallback) > 0 {
			return conv.resolveOutputName(prevOutputConfig.Fallback, prevOutput.Name())
		}
		return conv.resolveOutputName(conv.fallback(), prevOutput.Name())
	}
	return conv.resolveOutputName(output, prevOutput.Name())
}

func (conv *FlowConversation) fallback() string {
	fallback := conv.config.Fallback
	if len(fallback) > 0 {
		return fallback
	}
	return core.NotFoundOutputName
}

func (conv *FlowConversation) findPrevOutputConfig(prevOutput core.IOutput) flowOutput {
	if output, ok := conv.config.Outputs[prevOutput.Name()]; ok {
		return output
	}
	return flowOutput{empty: true}
}

func (conv *FlowConversation) findNextOutput(prevOutputName string) string {
	sequence := conv.config.Sequence
	for i := 0; i < len(sequence)-1; i++ {
		output := sequence[i]
		if prevOutputName == output {
			return sequence[i+1]
		}
	}
	return conv.fallback()
}

func (conv *FlowConversation) resolveOutputName(outputName, prevOutputName string) string {
	if outputName == core.NextOutputName {
		return conv.findNextOutput(prevOutputName)
	}
	return outputName
}
