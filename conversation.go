package ubahn

// IConversation describes a conversation that can be continued.
type IConversation interface {
	// Continue finds the next output to the given previous output and input.
	Continue(IInput, IConversationContext) IConversationContext

	// Empty returns true when the conversation is not initialized.
	Empty() bool

	// FlowName returns name of the current flow.
	FlowName() string
}

type nullConversation struct {
}

// Continue of the null conversation object returns blank output name.
func (conv *nullConversation) Continue(input IInput, ctx IConversationContext) IConversationContext {
	return NewConversationContext(conv, BlankOutput)
}

// Empty of the null conversation object returns true.
func (conv *nullConversation) Empty() bool {
	return true
}

// FlowName of the null conversation object returns empty string.
func (conv *nullConversation) FlowName() string {
	return ""
}

// NullConversation is a null object that implements IConversation interface.
var NullConversation = &nullConversation{}

// Conversation is the implementation of IConversation and may consist of multiple flows.
type Conversation struct {
	config           conversationConfig
	outputFactory    IOutputFactory
	flowConversation IConversation
}

// NewConversation creates a new instance of a conversation.
func NewConversation(file IConversationFile, outputFactory IOutputFactory) (IConversation, error) {
	conv := &Conversation{outputFactory: outputFactory}
	err := file.Parse(&conv.config)
	if err != nil {
		return NullConversation, err
	}
	return conv, nil
}

// NewConversationFromPath creates a new instance of a conversation,
// initialized from the given YAML file.
// If initialization filed, a null object is returned along with an error.
func NewConversationFromPath(conversationFilePath string, outputFactory IOutputFactory) (IConversation, error) {
	file, err := NewConversationFile(conversationFilePath)
	if err != nil {
		return NullConversation, err
	}

	return NewConversation(file, outputFactory)
}

func newContinuedConversation(
	config conversationConfig,
	flowConversation IConversation,
	outputFactory IOutputFactory) IConversation {
	return &Conversation{config: config, outputFactory: outputFactory, flowConversation: flowConversation}
}

// Continue finds the next output to the given previous output and input.
func (conv *Conversation) Continue(input IInput, ctx IConversationContext) IConversationContext {
	nextFlowName := conv.config.Triggers[input.Name()]
	if len(nextFlowName) == 0 {
		nextFlowName = conv.config.DefaultTrigger
	}

	// TODO
	return ctx
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *Conversation) Empty() bool {
	return false
}

// FlowName returns name of the current flow.
func (conv *Conversation) FlowName() string {
	return conv.flowConversation.FlowName()
}
