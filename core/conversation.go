package core

// IConversation describes a conversation that can be continued.
type IConversation interface {
	// Continue finds the next output to the given previous output and input.
	Continue(prevOutput IOutput, input IInput) string

	// Empty returns true when the conversation is not initialized.
	Empty() bool
}

type nullConversation struct {
}

func (conv *nullConversation) Continue(prevOutput IOutput, input IInput) string {
	return BlankOutputName
}

func (conv *nullConversation) Empty() bool {
	return true
}

// NullConversation is a null object that implements IConversation interface.
var NullConversation = &nullConversation{}
