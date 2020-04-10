package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NullConversation_Continue(t *testing.T) {
	input := NewNullInput("abc")
	prevOutput := BlankOutput
	nextOutput := NullConversation.Continue(prevOutput, input)

	assert.Equal(t, BlankOutputName, nextOutput.Name())
}

func Test_NullConversation_Empty(t *testing.T) {
	assert.True(t, NullConversation.Empty())
}
