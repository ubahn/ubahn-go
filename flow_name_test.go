package ubahn

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFlowName(t *testing.T) {
	name := newTestFlowName("weather/flows/city-weather.yml")

	assert.Equal(t, "city-weather", name)
}

func Test_NewFlowName_ManyDots(t *testing.T) {
	name := newTestFlowName("multiple.dots.file.yml")

	assert.Equal(t, "multiple.dots.file", name)
}

func Test_NewFlowName_NoDots(t *testing.T) {
	name := newTestFlowName("no-dots-file")

	assert.Equal(t, "no-dots-file", name)
}

func newTestFlowName(filePath string) string {
	file, _ := NewConversationFile(path.Join("./test_data/", filePath))
	return NewFlowName(file)
}
