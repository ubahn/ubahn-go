package ubahn

import (
	"fmt"
)

func newFakeInput(name string) IInput {
	return NewNullInput(name)
}

func newFakeOutput(name string) IOutput {
	return NewNullOutput(name)
}

func continueConversation(conv IConversation, prevOutput, input string) IConversationContext {
	inputContext := NewConversationContext(conv, newFakeOutput(prevOutput))
	return conv.Continue(newFakeInput(input), inputContext)
}

func newTestConversationFile(convName, testFileName string) (*ConversationFile, error) {
	path := fmt.Sprintf("./test_data/%s/flows/%s", convName, testFileName)
	return NewConversationFile(path)
}

func newTestFlowConversation(convName, testFileName string) (IConversation, error) {
	convFile, _ := newTestConversationFile(convName, testFileName)
	return NewFlowConversation(convFile, NewNullOutputFactory())
}

func startTestFlowConversation(convName, testFileName, startInput string) IConversationContext {
	conv, _ := newTestFlowConversation(convName, testFileName)
	inputContext := NewConversationContext(conv, BlankOutput)
	return conv.Continue(newFakeInput(startInput), inputContext)
}

func startTestFlowConversationDefault() IConversationContext {
	return startTestFlowConversation("weather", "city-weather.yml", "i-asks-city-weather")
}
