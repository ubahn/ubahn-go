package v2

import (
	"fmt"

	core "github.com/ubahn/ubahn-go/core"
)

// FlowConversation is the implementation of IConversation, specific to a flow.
// It's continued based on the sequence, defined in the given YAML file.
type FlowConversation struct {
	config        flowConfig
	outputFactory core.IOutputFactory
	flowName      string
}

// NewFlowConversation creates a new instance of a flow conversation.
func NewFlowConversation(file core.IConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	if file.Empty() {
		return core.NullConversation, fmt.Errorf("Main conversation file is not initialized")
	}

	conv := &FlowConversation{outputFactory: outputFactory, flowName: NewFlowName(file)}
	err := file.Parse(&conv.config)
	if err != nil {
		return core.NullConversation, err
	}
	return conv, nil
}

// Continue finds the next output to the given previous output and input.
// If no suitable output found, it returns blank output.
func (conv *FlowConversation) Continue(prevOutput core.IOutput, input core.IInput) (core.IOutput, core.IConversation) {
	nextOutputName := conv.matchOutput(prevOutput, input)
	return conv.outputFactory.Create(nextOutputName), conv
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *FlowConversation) Empty() bool {
	return false
}

// FlowName returns the name of the current flow.
func (conv *FlowConversation) FlowName() string {
	return conv.flowName
}

// matchOutput tries to next output name, considering previous output and user input.
func (conv *FlowConversation) matchOutput(prevOutput core.IOutput, input core.IInput) string {
	prevOutputConfig := conv.findPrevOutputConfig(prevOutput)
	if prevOutputConfig.empty {
		// If there was no previous output or we don't know it, we start with root output
		return conv.resolveRootOutput()
	}

	// Next output is inferred from the expected inputs of the previous output.
	// When input isn't expected, we return fallback: first we look for prev output fallback, then global.
	output := prevOutputConfig.ExpectedInputs[input.Name()]
	if len(output) == 0 {
		if len(prevOutputConfig.Fallback) > 0 {
			return prevOutputConfig.Fallback
		}
		return conv.fallback()
	}
	return output
}

// resolveRootOutput returns root output name if it exists, otherwise returns fallback.
func (conv *FlowConversation) resolveRootOutput() string {
	if len(conv.config.RootOutput) > 0 {
		return conv.config.RootOutput
	}
	return conv.fallback()
}

// fallback returns the main flow fallback if it exists, otherwise returns not found output.
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
