package ubahn

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func Test_NullOutputFactory_Create(t *testing.T) {
	factory := NewNullOutputFactory()
	outputName := fake.DomainName()
	output := factory.Create(outputName)

	assert.Equal(t, outputName, output.Name())
}
