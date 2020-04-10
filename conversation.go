package ubahn

import (
	"fmt"

	core "github.com/ubahn/ubahn-go/core"
	v1 "github.com/ubahn/ubahn-go/v1"
	v2 "github.com/ubahn/ubahn-go/v2"
)

// NewConversation creates a new instance of a conversation,
// initialized from the given YAML file.
// If initialization filed, a null object is returned along with an error.
func NewConversation(conversationFilePath string, outputFactory core.IOutputFactory) (core.IConversation, error) {
	file, err := core.NewConversationFile(conversationFilePath)
	if err != nil {
		return core.NullConversation, err
	}

	return newConversation(file, outputFactory)
}

func newConversation(file *core.ConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	switch file.Version {
	case 1:
		return newV1Conversation(file, outputFactory)
	case 2:
		return newV2Conversation(file, outputFactory)
	}

	return core.NullConversation, fmt.Errorf("Unsupported conversation file version %d", file.Version)
}

func newV1Conversation(file core.IConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	conversation, err := v1.NewConversation(file, outputFactory)
	if err != nil {
		return core.NullConversation, err
	}
	return conversation, nil
}

func newV2Conversation(file core.IConversationFile, outputFactory core.IOutputFactory) (core.IConversation, error) {
	conversation, err := v2.NewConversation(file, outputFactory)
	if err != nil {
		return core.NullConversation, err
	}
	return conversation, nil
}
