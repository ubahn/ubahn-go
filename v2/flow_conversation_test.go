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
	out, _ := startTestFlowConversationDefault()

	assert.Equal(t, "welcome", out.Name())
}

func Test_FlowConversation_Continue_StartWithoutRoot(t *testing.T) {
	out, _ := startTestFlowConversation("weather", "rootless.yml", "i-asks-rootless")

	// When root isn't specified, we use the first output.
	assert.Equal(t, "firstOutput", out.Name())
}

func Test_FlowConversation_Continue_ExpectedInput(t *testing.T) {
	prevOut, conv := startTestFlowConversationDefault()

	out, conv := conv.Continue(prevOut, newFakeInput("i-yes"))
	assert.Equal(t, "weather-report", out.Name())

	out, conv = conv.Continue(prevOut, newFakeInput("i-no"))
	assert.Equal(t, "bye", out.Name())
}

func Test_FlowConversation_Continue_LocalFallback(t *testing.T) {
	prevOut, conv := startTestFlowConversationDefault()

	out, conv := conv.Continue(prevOut, newFakeInput("i-blah"))

	assert.Equal(t, "welcome-clarification", out.Name())
}

func Test_FlowConversation_Continue_LastOutputFallback(t *testing.T) {
	prevOut, conv := startTestFlowConversationDefault()

	prevOut, conv = conv.Continue(prevOut, newFakeInput("i-yes"))
	out, conv := conv.Continue(prevOut, newFakeInput("i-blah"))

	assert.Equal(t, core.BlankOutputName, out.Name())
}

func Test_FlowConversation_Continue_Exit(t *testing.T) {
	prevOut, conv := startTestFlowConversationDefault()

	prevOut, conv = conv.Continue(prevOut, newFakeInput("i-no"))
	assert.Equal(t, "bye", prevOut.Name())

	out, conv := conv.Continue(prevOut, newFakeInput("i-asks-city-weather"))
	assert.Equal(t, "welcome", out.Name())
}

func Test_FlowConversation_Continue_GlobalFallback(t *testing.T) {
	prevOut, conv := startTestFlowConversationDefault()

	prevOut, conv = conv.Continue(prevOut, newFakeInput("i-maybe"))
	assert.Equal(t, "info", prevOut.Name())

	out, conv := conv.Continue(prevOut, newFakeInput("i-blah"))
	assert.Equal(t, "clarification", out.Name())
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

func startTestFlowConversationDefault() (core.IOutput, core.IConversation) {
	return startTestFlowConversation("weather", "city-weather.yml", "i-asks-city-weather")
}
