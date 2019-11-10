package ubahn

import (
	"fmt"

	core "github.com/ubahn/ubahn-go/core"
	v1 "github.com/ubahn/ubahn-go/v1"
)

// NewConversation creates a new instance of a conversation,
// initialized from the given YAML file.
// If initialization filed, a null object is returned along with an error.
func NewConversation(conversationFilePath string) (core.IConversation, error) {
	file, err := core.NewConversationFile(conversationFilePath)
	if err != nil {
		return core.NullConversation, err
	}

	return newConversation(file)
}

func newConversation(file *core.ConversationFile) (core.IConversation, error) {
	switch file.Version {
	case 1:
		return newV1Conversation(file)
	}

	return core.NullConversation, fmt.Errorf("Unsupported conversation file version %d", file.Version)
}

func newV1Conversation(file core.IConversationFile) (core.IConversation, error) {
	conversation, err := v1.NewConversation(file)
	if err != nil {
		return core.NullConversation, err
	}
	return conversation, nil
}
