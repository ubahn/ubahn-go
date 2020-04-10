package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

func Test_NewFlowConversation(t *testing.T) {
	conv, err := createFlowConversation("weather", "city-weather.yml")

	assert.False(t, conv.Empty())
	assert.Nil(t, err)
}

func Test_FlowConversation_Continue(t *testing.T) {

}

func Test_FlowConversation_Continue_Fallback(t *testing.T) {

}

func Test_FlowConversation_Continue_OutOfSequence(t *testing.T) {

}

func Test_FlowConversation_Continue_EmptySequence(t *testing.T) {

}

func Test_FlowConversation_Continue_EmptySequenceWithFallback(t *testing.T) {

}

func Test_FlowConversation_Continue_NotFound(t *testing.T) {

}

func startFlowConversation(convName, testFileName, trigger string) (core.IConversation, string) {
	conv, _ := createFlowConversation(convName, testFileName)
	nextOutputName := conv.Continue(core.BlankOutput, core.NewNullInput(trigger)).Name()
	return conv, nextOutputName
}

func createFlowConversation(convName, testFileName string) (core.IConversation, error) {
	path := fmt.Sprintf("../test_data/v2/%s/flows/%s", convName, testFileName)
	file, err := core.NewConversationFile(path)
	if err != nil {
		panic(err)
	}
	return NewFlowConversation(file, core.NewNullOutputFactory())
}
