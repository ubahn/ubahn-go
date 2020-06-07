package ubahn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFlowName(t *testing.T) {
	name := newTestFlowName("../test_data/v2/weather/flows/city-weather.yml")

	assert.Equal(t, "city-weather", name)
}

func Test_NewFlowName_ManyDots(t *testing.T) {
	name := newTestFlowName("../test_data/v2/multiple.dots.file.yml")

	assert.Equal(t, "multiple.dots.file", name)
}

func Test_NewFlowName_NoDots(t *testing.T) {
	name := newTestFlowName("../test_data/v2/no-dots-file")

	assert.Equal(t, "no-dots-file", name)
}

func newTestFlowName(path string) string {
	file, _ := NewConversationFile(path)
	return NewFlowName(file)
}
