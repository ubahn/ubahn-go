package ubahn

import (
	"path"
	"path/filepath"
)

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
	config        conversationConfig
	outputFactory IOutputFactory
	rootPath      string
}

// RestoreConversation creates an instance of either flow or root conversation, based on input parameters.
// This function is useful when previous conversation state was stored in a session and needs to be restored later.
func RestoreConversation(rootPath, flowName string, outputFactory IOutputFactory) (IConversation, error) {
	rootConv, err := NewConversationFromPath(path.Join(rootPath, "conversation.yml"), outputFactory)
	if err != nil || len(flowName) == 0 {
		return rootConv, err
	}

	return (rootConv.(*Conversation)).newFlowConversation(flowName)
}

// NewConversation creates a new instance of a conversation.
func NewConversation(file IConversationFile, outputFactory IOutputFactory) (IConversation, error) {
	conv := &Conversation{outputFactory: outputFactory, rootPath: file.FilePath()}
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

// Continue finds the next output to the given previous output and input.
func (conv *Conversation) Continue(input IInput, ctx IConversationContext) IConversationContext {
	nextFlowName := conv.config.Triggers[input.Name()]
	if len(nextFlowName) == 0 {
		if len(conv.config.DefaultTrigger) == 0 {
			return NewConversationContext(conv, BlankOutput)
		}
		nextFlowName = conv.config.DefaultTrigger
	}
	flowConv, err := conv.newFlowConversation(nextFlowName)
	if err != nil {
		panic(err)
	}
	return flowConv.Continue(input, ctx)
}

// Empty returns true when the conversation is not initialized.
// This implementation is considered to be always initialized.
func (conv *Conversation) Empty() bool {
	return false
}

// FlowName returns empty string because this type of conversation doesn’t have a flow.
func (conv *Conversation) FlowName() string {
	return ""
}

func (conv *Conversation) newFlowConversationFilePath(flowName string) string {
	return path.Join(
		filepath.Dir(conv.rootPath),
		"flows",
		flowName+".yml")
}

func (conv *Conversation) newFlowConversationFile(flowName string) IConversationFile {
	file, err := NewConversationFile(conv.newFlowConversationFilePath(flowName))
	if err != nil {
		panic(err)
	}
	return file
}

func (conv *Conversation) newFlowConversation(flowName string) (IConversation, error) {
	file := conv.newFlowConversationFile(flowName)
	return NewFlowConversation(file, conv.outputFactory, conv)
}
