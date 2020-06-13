package ubahn

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	ctx := startTestFlowConversationDefault()

	assert.Equal(t, "welcome", ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_StartWithoutRoot(t *testing.T) {
	ctx := startTestFlowConversation("weather", "rootless.yml", "i-asks-rootless")

	// When root isn't specified, we use the first output.
	assert.Equal(t, "firstOutput", ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_ExpectedInput(t *testing.T) {
	originalCtx := startTestFlowConversationDefault()

	ctx := originalCtx.Conversation().Continue(newFakeInput("i-yes"), originalCtx)
	assert.Equal(t, "weather-report", ctx.LastOutput().Name())

	ctx = originalCtx.Conversation().Continue(newFakeInput("i-no"), originalCtx)
	assert.Equal(t, "bye", ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_LocalFallback(t *testing.T) {
	ctx := startTestFlowConversationDefault()

	ctx = ctx.Conversation().Continue(newFakeInput("i-blah"), ctx)

	assert.Equal(t, "welcome-clarification", ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_LastOutputFallback(t *testing.T) {
	ctx := startTestFlowConversationDefault()

	ctx = ctx.Conversation().Continue(newFakeInput("i-yes"), ctx)
	ctx = ctx.Conversation().Continue(newFakeInput("i-blah"), ctx)

	assert.Equal(t, BlankOutputName, ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_Exit(t *testing.T) {
	ctx := startTestFlowConversationDefault()

	ctx = ctx.Conversation().Continue(newFakeInput("i-no"), ctx)
	assert.Equal(t, "bye", ctx.LastOutput().Name())

	ctx = ctx.Conversation().Continue(newFakeInput("i-asks-city-weather"), ctx)
	assert.Equal(t, "welcome", ctx.LastOutput().Name())
}

func Test_FlowConversation_Continue_GlobalFallback(t *testing.T) {
	ctx := startTestFlowConversationDefault()

	ctx = ctx.Conversation().Continue(newFakeInput("i-maybe"), ctx)
	assert.Equal(t, "info", ctx.LastOutput().Name())

	ctx = ctx.Conversation().Continue(newFakeInput("i-blah"), ctx)
	assert.Equal(t, "clarification", ctx.LastOutput().Name())
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
