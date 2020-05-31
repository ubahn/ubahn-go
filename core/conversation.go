package core

// IConversation describes a conversation that can be continued.
type IConversation interface {
	// Continue finds the next output to the given previous output and input.
	Continue(prevOutput IOutput, input IInput) (IOutput, IConversation)

	// Empty returns true when the conversation is not initialized.
	Empty() bool
}

type nullConversation struct {
}

// Continue of the null conversation object returns blank output name.
func (conv *nullConversation) Continue(prevOutput IOutput, input IInput) (IOutput, IConversation) {
	return BlankOutput, conv
}

// Empty of the null conversatino object returns true.
func (conv *nullConversation) Empty() bool {
	return true
}

// NullConversation is a null object that implements IConversation interface.
var NullConversation = &nullConversation{}
