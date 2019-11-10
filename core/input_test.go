package core

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestNullInput(t *testing.T) {
	name := fake.Model()
	input := NewNullInput(name)

	assert.Equal(t, name, input.Name())
}
