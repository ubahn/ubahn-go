package v2

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

}

func Test_Continue_Fallback(t *testing.T) {

}

func Test_Continue_OutOfSequence(t *testing.T) {

}

func Test_Continue_EmptySequence(t *testing.T) {

}

func Test_Continue_EmptySequenceWithFallback(t *testing.T) {

}

func Test_Continue_NotFound(t *testing.T) {

}

func startConversation(testFileName, trigger string) (core.IConversation, string) {
	conv, _ := createConversation(testFileName)
	nextOutputName := conv.Continue(core.BlankOutput, core.NewNullInput(trigger)).Name()
	return conv, nextOutputName
}

func newFakeInput(name string) core.IInput {
	return core.NewNullInput(name)
}

func newFakeOutput(name string) core.IOutput {
	return core.NewNullOutput(name)
}

func continueConversation(conv core.IConversation, prevOutput, input string) string {
	return conv.Continue(newFakeOutput(prevOutput), newFakeInput(input)).Name()
}

func createConversation(testFileName string) (core.IConversation, error) {
	path := fmt.Sprintf("../test_data/v1/%s", testFileName)
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewConversation(file, core.NewNullOutputFactory())
}
