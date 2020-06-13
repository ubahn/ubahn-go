package ubahn

import (
	"fmt"
	"path"
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
	return NewConversationFile(newTestConversationFilePath(convName, testFileName))
}

func newTestConversationFilePath(convName, testFileName string) string {
	return fmt.Sprintf("./test_data/%s/flows/%s", convName, testFileName)
}

func newTestFlowConversation(convName, testFileName string) (IConversation, error) {
	convFile, _ := newTestConversationFile(convName, testFileName)
	rootConvFile, err := NewConversationFile(path.Join("./test_data", convName, "conversation.yml"))
	if err != nil {
		panic(err)
	}
	outputFactory := NewNullOutputFactory()
	rootConv, err := NewConversation(rootConvFile, outputFactory)
	if err != nil {
		panic(err)
	}
	return NewFlowConversation(convFile, outputFactory, rootConv)
}

func startTestFlowConversation(convName, testFileName, startInput string) IConversationContext {
	conv, _ := newTestFlowConversation(convName, testFileName)
	inputContext := NewConversationContext(conv, BlankOutput)
	return conv.Continue(newFakeInput(startInput), inputContext)
}

func startTestFlowConversationDefault() IConversationContext {
	return startTestFlowConversation("weather", "city-weather.yml", "i-asks-city-weather")
}
