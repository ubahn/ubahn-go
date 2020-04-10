package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

func Test_NewConversation(t *testing.T) {
	conv, err := createConversation("weather")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)
}

func createConversation(convName string) (core.IConversation, error) {
	path := fmt.Sprintf("../test_data/v2/%s/conversation.yml", convName)
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewConversation(file, core.NewNullOutputFactory())
}
