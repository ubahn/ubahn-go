package ubahn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConversationContext(t *testing.T) {
	prevOutput, conv := startTestFlowConversationDefault()
	input := newFakeInput("i-smth")
	nextOutput := newFakeOutput("blah-blah")
	history := make([]IHistoryItem, 1)
	ctx := NewConversationContext(conv, prevOutput, input, nextOutput, history)

	assert.Equal(t, prevOutput.Name(), ctx.PrevOutput().Name())
	assert.Equal(t, input.Name(), ctx.Input().Name())
	assert.Equal(t, nextOutput.Name(), ctx.NextOutput().Name())
	assert.True(t, ctx.HasHistory())
}

func Test_NewEmptyConversationContext(t *testing.T) {
	_, conv := startTestFlowConversationDefault()
	input := newFakeInput("i-smth")
	ctx := NewEmptyConversationContext(conv, input)

	assert.Equal(t, BlankOutputName, ctx.PrevOutput().Name())
	assert.Equal(t, input.Name(), ctx.Input().Name())
	assert.Equal(t, BlankOutputName, ctx.NextOutput().Name())
	assert.False(t, ctx.HasHistory())
}

func Test_Next(t *testing.T) {
	// input1 := newFakeInput("i-user-welcome")
	// conv, _ := newTestFlowConversation("weather", "city-weather.yml")
	// nextOutput, conv := conv.Continue(BlankOutput, input1)

	// ctx := NewConversationContext(conv, prevOutput, input1)
	// input2 := newFakeInput("i-yes")
	// nextOutput, conv := conv.Continue(prevOutput, input2)

	// newContext := ctx.Next(input2, nextOutput, conv)

	// assert.True(t, newContext.HasHistory())
	// lastHistoryItem := newContext.History()[0]
	// assert.Equal(t, input1.Name(), lastHistoryItem.InputName)
	// assert.Equal(t, prevOutput.Name(), lastHistoryItem.OutputName)
	// assert.Equal(t, conv.FlowName(), lastHistoryItem.FlowName)
}
