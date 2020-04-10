package v2

import (
	core "github.com/ubahn/ubahn-go/core"
)

// Conversation is the implementation of IConversation and may consist of multiple flows.
type Conversation struct {
	config        conversationConfig
	outputFactory core.IOutputFactory
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

// Continue finds the next output to the given previous output and input.
func (conv *Conversation) Continue(prevOutput core.IOutput, input core.IInput) core.IOutput {
	return core.BlankOutput
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *Conversation) Empty() bool {
	return false
}
