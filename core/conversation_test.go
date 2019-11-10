package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullConversationContinue(t *testing.T) {
	input := NewNullInput("abc")
	prevOutput := BlankOutput
	nextOutputName := NullConversation.Continue(prevOutput, input)

	assert.Equal(t, BlankOutputName, nextOutputName)
}

func TestNullConversationEmpty(t *testing.T) {
	assert.True(t, NullConversation.Empty())
}
