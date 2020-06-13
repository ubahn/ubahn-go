package ubahn

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func Test_NullOutput(t *testing.T) {
	name := fake.Model()
	output := NewNullOutput(name)

	assert.Equal(t, name, output.Name())

	res, err := output.Send()
	assert.Nil(t, res)
	assert.Nil(t, err)
}

func Test_BlankOutput(t *testing.T) {
	assert.NotNil(t, BlankOutput)
	assert.NotEmpty(t, BlankOutput.Name())
}
