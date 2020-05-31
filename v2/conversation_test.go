package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

func Test_NewConversation(t *testing.T) {
	assertConversationCreated(t, "weather")
}

func Test_NewConversation_Empty(t *testing.T) {
	assertConversationCreated(t, "empty")
}

func Test_NewConversation_WeirdYaml(t *testing.T) {
	assertConversationCreated(t, "weird")
}

func Test_Continue_Triggers(t *testing.T) {
	conv, _ := createConversation("weather")

	nextOut, nextConv := continueConversation(conv, core.BlankOutputName, "i-user-says-welcome")
	assert.Equal(t, "welcome", nextOut)

	flowConv, ok := nextConv.(*FlowConversation)
	assert.True(t, ok)
	assert.NotNil(t, flowConv)
	assert.Equal(t, "city-weather", flowConv.FlowName())
}

func assertConversationCreated(t *testing.T, convName string) core.IConversation {
	conv, err := createConversation("empty")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)

	return conv
}

func createConversation(convName string) (core.IConversation, error) {
	path := fmt.Sprintf("../test_data/v2/%s/conversation.yml", convName)
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewConversation(file, core.NewNullOutputFactory())
}
