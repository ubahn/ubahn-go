package v1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

func Test_NewConversation(t *testing.T) {
	conv, err := createConversation("weather.yml")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)
}

func Test_Continue(t *testing.T) {
	conv, outputName := startConversation("weather.yml", "i-user-welcome")
	assert.Equal(t, "welcome", outputName)

	outputName = continueConversation(conv, outputName, "i-yes")
	assert.Equal(t, "weather-report", outputName)

	outputName = continueConversation(conv, outputName, "")
	assert.Equal(t, "clarification", outputName)

	outputName = continueConversation(conv, "welcome", "i-no")
	assert.Equal(t, "bye", outputName)
}

func Test_Continue_Fallback(t *testing.T) {
	conv, outputName := startConversation("fallbacks.yml", "i-any")
	assert.Equal(t, "a", outputName)

	assert.Equal(t, "b", continueConversation(conv, outputName, "i-yes"))
	assert.Equal(t, "c", continueConversation(conv, outputName, "i-no"))
	assert.Equal(t, "b", continueConversation(conv, outputName, "i-unknown"))

	assert.Equal(t, "d", continueConversation(conv, "b", "i-unknown"))
}

func Test_Continue_OutOfSequence(t *testing.T) {
	conv, outputName := startConversation("out_of_sequence.yml", "i-any")
	assert.Equal(t, "step1", outputName)

	assert.Equal(t, "step2", continueConversation(conv, outputName, "i-yes"))
	assert.Equal(t, "end", continueConversation(conv, outputName, "i-no"))
	assert.Equal(t, "unexpected1", continueConversation(conv, outputName, "i-unknown"))

	assert.Equal(t, "step2", continueConversation(conv, "unexpected1", "i-yes"))
	assert.Equal(t, "end", continueConversation(conv, "unexpected1", "i-no"))
	assert.Equal(t, "unexpected2", continueConversation(conv, "unexpected1", "i-blahblah"))
}

func Test_Continue_EmptySequence(t *testing.T) {
	_, nextOutputName := startConversation("weather_empty.yml", "i-user-welcome")

	assert.Equal(t, core.NotFoundOutputName, nextOutputName)
}

func Test_Continue_EmptySequenceWithFallback(t *testing.T) {
	_, nextOutputName := startConversation("weather_empty_fallback.yml", "i-user-welcome")

	assert.Equal(t, "clarification", nextOutputName)
}

func Test_Continue_NotFound(t *testing.T) {
	conv, outputName := startConversation("no_fallback.yml", "i-any")

	assert.Equal(t, "a", outputName)

	assert.Equal(t, "b", continueConversation(conv, outputName, "i-yes"))
	assert.Equal(t, core.NotFoundOutputName, continueConversation(conv, "b", "i-yes"))

	assert.Equal(t, core.NotFoundOutputName, continueConversation(conv, outputName, "i-blah-blah"))
}

func startConversation(testFileName, trigger string) (core.IConversation, string) {
	conv, _ := createConversation(testFileName)
	nextOutputName := conv.Continue(core.BlankOutput, core.NewNullInput(trigger))
	return conv, nextOutputName
}

func newFakeInput(name string) core.IInput {
	return core.NewNullInput(name)
}

func newFakeOutput(name string) core.IOutput {
	return core.NewNullOutput(name)
}

func continueConversation(conv core.IConversation, prevOutput, input string) string {
	return conv.Continue(newFakeOutput(prevOutput), newFakeInput(input))
}

func createConversation(testFileName string) (core.IConversation, error) {
	path := fmt.Sprintf("../test_data/v1/%s", testFileName)
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewConversation(file)
}
