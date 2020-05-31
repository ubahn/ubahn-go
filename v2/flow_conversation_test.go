package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/ubahn/ubahn-go/core"
)

func Test_NewFlowConversation(t *testing.T) {
	assertInitializedFlowConversation(t, "weather", "city-weather.yml")
}

func Test_NewFlowConversation_InvalidConvFile(t *testing.T) {
	assertNotInitializedFlowConversation(t, "weird", "invalid.yml")
}

func Test_NewFlowConversation_InvalidFormat(t *testing.T) {
	assertNotInitializedFlowConversation(t, "weather", "invalid.yml")
}

func Test_FlowConversation_Continue_StartWithRoot(t *testing.T) {
	out, _ := startTestFlowConversation("weather", "city-weather.yml", "i-asks-city-weather")

	assert.Equal(t, "welcome", out.Name())
}

func Test_FlowConversation_Continue_Fallback(t *testing.T) {

}

func Test_FlowConversation_Continue_NotFound(t *testing.T) {

}

func newTestConversationFile(convName, testFileName string) (*core.ConversationFile, error) {
	path := fmt.Sprintf("../test_data/v2/%s/flows/%s", convName, testFileName)
	return core.NewConversationFile(path)
}

func newTestFlowConversation(convName, testFileName string) (core.IConversation, error) {
	convFile, _ := newTestConversationFile(convName, testFileName)
	return NewFlowConversation(convFile, core.NewNullOutputFactory())
}

func assertInitializedFlowConversation(t *testing.T, convName, testFileName string) {
	conv, err := newTestFlowConversation(convName, testFileName)

	assert.Falsef(t, conv.Empty(), "Expected flow conversation not to be empty")
	assert.Nilf(t, err,
		"Expected flow conversation to have been initialized without error, but there's error: %v", err)
}

func assertNotInitializedFlowConversation(t *testing.T, convName, testFileName string) {
	conv, err := newTestFlowConversation(convName, testFileName)

	assert.Truef(t, conv.Empty(), "Expected flow conversation to be empty, but it wasn't")
	assert.NotNilf(t, err,
		"Expected flow conversation to have been initialized with error, but there was no error")
}

func startTestFlowConversation(convName, testFileName, startInput string) (core.IOutput, core.IConversation) {
	conv, _ := newTestFlowConversation(convName, testFileName)
	return conv.Continue(core.BlankOutput, core.NewResolvedInput(startInput, nil))
}
