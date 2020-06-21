package ubahn

import (
	"fmt"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NullConversation_Continue(t *testing.T) {
	ctx := NewEmptyConversationContext(NullConversation)
	ctx = NullConversation.Continue(newFakeInput("fake"), ctx)

	assert.Equal(t, BlankOutputName, ctx.LastOutput().Name())
}

func Test_NullConversation_Empty(t *testing.T) {
	assert.True(t, NullConversation.Empty())
}

func Test_RestoreConversation_Root(t *testing.T) {
	path := testConversationPath("weather")
	conv, err := RestoreConversation(path, "", NewNullOutputFactory())

	assert.Nil(t, err)
	assert.False(t, conv.Empty())
	_, ok := conv.(*Conversation)
	assert.True(t, ok)
}

func Test_RestoreConversation_Flow(t *testing.T) {
	path := testConversationPath("weather")
	flowName := "city-weather"
	conv, err := RestoreConversation(path, flowName, NewNullOutputFactory())

	assert.Nil(t, err)
	assert.False(t, conv.Empty())
	_, ok := conv.(*FlowConversation)
	assert.True(t, ok)
	assert.Equal(t, flowName, conv.FlowName())
}

func Test_RestoreConversation_NotExists(t *testing.T) {
	path := testConversationPath("blah")
	conv, err := RestoreConversation(path, "yoyo", NewNullOutputFactory())

	assert.NotNil(t, err)
	assert.True(t, conv.Empty())
}

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

	ctx := continueConversation(conv, BlankOutputName, "i-asks-city-weather")
	assert.Equal(t, "welcome", ctx.LastOutput().Name())

	flowConv, ok := ctx.Conversation().(*FlowConversation)
	assert.True(t, ok)
	assert.NotNil(t, flowConv)
	assert.Equal(t, "city-weather", flowConv.FlowName())
}

func assertConversationCreated(t *testing.T, convName string) IConversation {
	conv, err := createConversation("empty")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)

	return conv
}

func createConversation(convName string) (IConversation, error) {
	path := path.Join(testConversationPath(convName), "conversation.yml")
	return NewConversationFromPath(path, NewNullOutputFactory())
}

func testConversationPath(convName string) string {
	return fmt.Sprintf("./test_data/%s/", convName)
}
