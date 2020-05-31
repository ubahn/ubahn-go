package v2

import (
	core "github.com/ubahn/ubahn-go/core"
)

func newFakeInput(name string) core.IInput {
	return core.NewNullInput(name)
}

func newFakeOutput(name string) core.IOutput {
	return core.NewNullOutput(name)
}

func continueConversation(conv core.IConversation, prevOutput, input string) (string, core.IConversation) {
	nextOutput, nextConv := conv.Continue(newFakeOutput(prevOutput), newFakeInput(input))
	return nextOutput.Name(), nextConv
}
