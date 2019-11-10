package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

type fakeInput struct {
	name string
}

func (in *fakeInput) Name() string {
	return in.name
}

func TestNewConversation(t *testing.T) {
	conv, err := createConversation("../test_data/weather.yml")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)
}

func TestContinue(t *testing.T) {
	conv, nextOutputName := startConversation("../test_data/weather.yml", "i-user-welcome")
	assert.Equal(t, "welcome", nextOutputName)

	nextOutputName = continueConversation(conv, nextOutputName, "i-yes")
	assert.Equal(t, "weather-report", nextOutputName)

	nextOutputName = continueConversation(conv, nextOutputName, "")
	assert.Equal(t, "clarification", nextOutputName)

	nextOutputName = continueConversation(conv, "welcome", "i-no")
	assert.Equal(t, "bye", nextOutputName)
}

func TestContinue_EmptySequence(t *testing.T) {
	_, nextOutputName := startConversation("../test_data/weather_empty.yml", "i-user-welcome")

	assert.Equal(t, core.BlankOutputName, nextOutputName)
}

func TestContinue_EmptySequenceWithFallback(t *testing.T) {
	_, nextOutputName := startConversation("../test_data/weather_empty_fallback.yml", "i-user-welcome")

	assert.Equal(t, "clarification", nextOutputName)
}

func startConversation(path, trigger string) (core.IConversation, string) {
	conv, _ := createConversation(path)
	nextOutputName := conv.Continue(core.BlankOutput, &fakeInput{name: trigger})
	return conv, nextOutputName
}

func newFakeInput(name string) core.IInput {
	return &fakeInput{name: name}
}

func newFakeOutput(name string) core.IOutput {
	return core.NewNullOutput(name)
}

func continueConversation(conv core.IConversation, prevOutput, input string) string {
	return conv.Continue(newFakeOutput(prevOutput), newFakeInput(input))
}

func createConversation(path string) (core.IConversation, error) {
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewConversation(file)
}
