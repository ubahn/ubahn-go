package ubahn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewHistoryItem(t *testing.T) {
	input := newFakeInput("i-fake")
	nextOutput := newFakeOutput("fake-out")
	conv, _ := newTestFlowConversation("weather", "city-weather.yml")
	item := NewHistoryItem(input, nextOutput, conv)

	assert.Equal(t, input.Name(), item.InputName())
	assert.Equal(t, nextOutput.Name(), item.OutputName())
	assert.Equal(t, conv.FlowName(), item.FlowName())
	assert.True(t, item.Timestamp().Unix() >= time.Now().UTC().Unix())
	assert.True(t, len(item.String()) > 0)
}
