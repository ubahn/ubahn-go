package core

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func Test_Name(t *testing.T) {
	name := fake.DomainName()
	input := NewResolvedInput(name, nil)

	assert.Equal(t, name, input.Name())
}

func Test_Data(t *testing.T) {
	data := []string{"hey"}
	input := NewResolvedInput("", data)

	assert.Equal(t, data, input.Data())
}

func Test_IsResolved(t *testing.T) {
	assert.True(t, NewResolvedInput("i-test", nil).IsResolved())
	assert.False(t, NewResolvedInput("", nil).IsResolved())
}
