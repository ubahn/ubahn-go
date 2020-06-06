package v2

import (
	core "github.com/ubahn/ubahn-go/core"
)

// Conversation is the implementation of IConversation and may consist of multiple flows.
type Conversation struct {
	config           conversationConfig
	outputFactory    core.IOutputFactory
	flowConversation core.IConversation
}

// NewConversation creates a new instance of a conversation.
func NewConversation(file core.IConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	conv := &Conversation{outputFactory: outputFactory}
	err := file.Parse(&conv.config)
	if err != nil {
		return core.NullConversation, err
	}
	return conv, nil
}

func newContinuedConversation(
	config conversationConfig,
	flowConversation core.IConversation,
	outputFactory core.IOutputFactory) core.IConversation {
	return &Conversation{config: config, outputFactory: outputFactory, flowConversation: flowConversation}
}

// Continue finds the next output to the given previous output and input.
func (conv *Conversation) Continue(prevOutput core.IOutput, input core.IInput) (core.IOutput, core.IConversation) {
	nextFlowName := conv.config.Triggers[input.Name()]
	if len(nextFlowName) == 0 {
		nextFlowName = conv.config.DefaultTrigger
	}

	return core.BlankOutput, conv
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *Conversation) Empty() bool {
	return false
}
