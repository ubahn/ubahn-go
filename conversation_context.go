package ubahn

import (
	"time"
)

// IConversationContext describes state of a conversation, including history of inputs and outputs,
// latest flow, output and input.
type IConversationContext interface {
	Conversation() IConversation
	PrevOutput() IOutput
	Input() IInput
	NextOutput() IOutput
	History() []IHistoryItem
	HasHistory() bool
	Next(nextInput IInput, nextOutput IOutput, nextConversation IConversation) IConversationContext
}

// IHistoryItem describes historical piece of a conversation, where you can see what was the input, the output,
// the flow and when it was recorded.
type IHistoryItem interface {
	InputName() string
	OutputName() string
	FlowName() string
	Timestamp() time.Time
	String() string
}

// ConversationContext contains state of a conversation, including history of inputs and outputs,
// latest flow, output and input.
type ConversationContext struct {
	conversation IConversation
	prevOutput   IOutput
	input        IInput
	nextOutput   IOutput
	history      []IHistoryItem
}

// NewConversationContext creates a new instance of a conversation context.
func NewConversationContext(
	conversation IConversation,
	prevOutput IOutput,
	input IInput,
	nextOutput IOutput,
	history []IHistoryItem) IConversationContext {
	return &ConversationContext{
		conversation: conversation,
		prevOutput:   prevOutput,
		input:        input,
		nextOutput:   nextOutput,
		history:      history}
}

// NewEmptyConversationContext creates a new instance of a conversation context with empty history and outputs.
func NewEmptyConversationContext(
	conversation IConversation,
	input IInput) IConversationContext {
	return NewConversationContext(
		conversation,
		BlankOutput,
		input,
		BlankOutput,
		make([]IHistoryItem, 0))
}

// Conversation returns current conversation.
func (ctx *ConversationContext) Conversation() IConversation {
	return ctx.conversation
}

// PrevOutput returns previous output.
func (ctx *ConversationContext) PrevOutput() IOutput {
	return ctx.prevOutput
}

// Input returns last user input.
func (ctx *ConversationContext) Input() IInput {
	return ctx.input
}

// NextOutput returns the next output.
func (ctx *ConversationContext) NextOutput() IOutput {
	return ctx.nextOutput
}

// History returns conversation history.
func (ctx *ConversationContext) History() []IHistoryItem {
	return ctx.history
}

// HasHistory returns true if there's at least one item in the history.
func (ctx *ConversationContext) HasHistory() bool {
	return len(ctx.History()) > 0
}

// Next creates a copy of this context, with new output, and stores previous state to history.
func (ctx *ConversationContext) Next(
	nextInput IInput,
	nextOutput IOutput,
	nextConversation IConversation) IConversationContext {
	historyItem := NewHistoryItem(ctx.Input(), ctx.NextOutput(), ctx.Conversation())
	newHistory := append(ctx.History(), historyItem)
	return NewConversationContext(nextConversation, ctx.NextOutput(), nextInput, nextOutput, newHistory)
}
