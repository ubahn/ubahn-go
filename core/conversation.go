package core

// IConversation defines a conversation that can be continued.
type IConversation interface {
	// Continue finds the next output to the given previous output and input.
	Continue(prevOutput IOutput, input IInput) string

	// Empty returns true when the conversation is not initialized.
	Empty() bool
}

// emptyConversation is a null object which implements empty IConversation.
type emptyConversation struct {
}

func (conv *emptyConversation) Continue(prevOutput IOutput, input IInput) string {
	return BlankOutputName
}

func (conv *emptyConversation) Empty() bool {
	return true
}

var NullConversation = &emptyConversation{}
