package ubahn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeInput struct {
	name string
}

func (in *fakeInput) Name() string {
	return in.name
}

func TestNewConversation_InvalidFile(t *testing.T) {
	conv, err := NewConversation("")
	assertFailure := func() {
		assert.True(t, conv.Empty)
		assert.NotNil(t, err)
	}
	assertFailure()

	conv, err = NewConversation("test_data/invalid.yml")
	assertFailure()
}

func TestNewConversation(t *testing.T) {
	conv, err := NewConversation("test_data/weather.yml")

	assert.False(t, conv.Empty)
	assert.Nil(t, err)
}

func TestContinue(t *testing.T) {
	conv, nextOutputName := startConversation("test_data/weather.yml", "i-user-welcome")
	assert.Equal(t, "welcome", nextOutputName)

	nextOutputName = continueConversation(conv, nextOutputName, "i-yes")
	assert.Equal(t, "weather-report", nextOutputName)

	nextOutputName = continueConversation(conv, nextOutputName, "")
	assert.Equal(t, "clarification", nextOutputName)

	nextOutputName = continueConversation(conv, "welcome", "i-no")
	assert.Equal(t, "bye", nextOutputName)
}

func TestContinue_EmptySequence(t *testing.T) {
	_, nextOutputName := startConversation("test_data/weather_empty.yml", "i-user-welcome")

	assert.Equal(t, BlankOutputName, nextOutputName)
}

func TestContinue_EmptySequenceWithFallback(t *testing.T) {
	_, nextOutputName := startConversation("test_data/weather_empty_fallback.yml", "i-user-welcome")

	assert.Equal(t, "clarification", nextOutputName)
}

func startConversation(path, trigger string) (*Conversation, string) {
	conv, _ := NewConversation(path)
	nextOutputName := conv.Continue(BlankOutput, &fakeInput{name: trigger})
	return conv, nextOutputName
}

func newFakeInput(name string) IInput {
	return &fakeInput{name: name}
}

func newFakeOutput(name string) IOutput {
	return &nullOutput{name: name}
}

func continueConversation(conv IConversation, prevOutput, input string) string {
	return conv.Continue(newFakeOutput(prevOutput), newFakeInput(input))
}
