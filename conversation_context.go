package ubahn

// IConversationContext describes state of a conversation.
type IConversationContext interface {
	Conversation() IConversation
	LastOutput() IOutput
}

// ConversationContext implements state of a conversation.
type ConversationContext struct {
	conversation IConversation
	lastOutput   IOutput
}

// NewConversationContext creates a new instance of a conversation context.
func NewConversationContext(
	conv IConversation,
	lastOutput IOutput) IConversationContext {
	return &ConversationContext{
		conversation: conv,
		lastOutput:   lastOutput}
}

// NewEmptyConversationContext creates a new instance of a conversation context with empty output.
func NewEmptyConversationContext(conv IConversation) IConversationContext {
	return NewConversationContext(conv, BlankOutput)
}

// Conversation returns current conversation.
func (ctx *ConversationContext) Conversation() IConversation {
	return ctx.conversation
}

// LastOutput returns last known output.
func (ctx *ConversationContext) LastOutput() IOutput {
	return ctx.lastOutput
}
