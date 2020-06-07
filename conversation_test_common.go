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

func continueConversation(conv IConversation, prevOutput, input string) (string, IConversation) {
	nextOutput, nextConv := conv.Continue(newFakeOutput(prevOutput), newFakeInput(input))
	return nextOutput.Name(), nextConv
}

func newTestConversationFile(convName, testFileName string) (*ConversationFile, error) {
	path := fmt.Sprintf("./test_data/%s/flows/%s", convName, testFileName)
	return NewConversationFile(path)
}

func newTestFlowConversation(convName, testFileName string) (IConversation, error) {
	convFile, _ := newTestConversationFile(convName, testFileName)
	return NewFlowConversation(convFile, NewNullOutputFactory())
}

func startTestFlowConversation(convName, testFileName, startInput string) (IOutput, IConversation) {
	conv, _ := newTestFlowConversation(convName, testFileName)
	return conv.Continue(BlankOutput, NewResolvedInput(startInput, nil))
}

func startTestFlowConversationDefault() (IOutput, IConversation) {
	return startTestFlowConversation("weather", "city-weather.yml", "i-asks-city-weather")
}
