package ubahn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConversationFile(t *testing.T) {
	file, err := newConversationFile("test_data/weather.yml")

	assert.False(t, file.empty)
	assert.Nil(t, err)
	assert.NotEmpty(t, file.data)
	assert.Equal(t, 1, file.version)
}

func TestNewConversationFile_InvalidFile(t *testing.T) {
	file, err := newConversationFile("test_data/invalid.yml")

	assert.True(t, file.empty)
	assert.NotNil(t, err)
}

func TestNewConversationFile_WhenFileDoesNotExist(t *testing.T) {
	file, err := newConversationFile("")

	assert.True(t, file.empty)
	assert.NotNil(t, err)
}

func TestNewConversationFile_Version(t *testing.T) {
	file, _ := newConversationFile("test_data/v1.yml")

	assert.Equal(t, 1, file.version)

	file, _ = newConversationFile("test_data/v2.yml")

	assert.Equal(t, 2, file.version)
}

func TestV1Config(t *testing.T) {
	file, _ := newConversationFile("test_data/weather.yml")
	config, err := file.V1Config()

	assert.Nil(t, err)

	assert.Equal(t, 2, len(config.Sequence))
	assert.Equal(t, "welcome", config.Sequence[0])
	assert.Equal(t, "weather-report", config.Sequence[1])

	assert.Equal(t, 1, len(config.Triggers))
	assert.Equal(t, "i-user-welcome", config.Triggers[0])

	assert.Equal(t, "clarification", config.Fallback)

	assert.Equal(t, 1, len(config.Outputs))

	welcomeOutput := config.Outputs["welcome"]
	assert.NotNil(t, welcomeOutput)
	assert.Equal(t, 2, len(welcomeOutput.ExpectedInputs))
	assert.Equal(t, "next", welcomeOutput.ExpectedInputs["i-yes"])
	assert.Equal(t, "bye", welcomeOutput.ExpectedInputs["i-no"])
	assert.Equal(t, "welcome-clarification", welcomeOutput.Fallback)
}

func TestV1Config_WrongVersion(t *testing.T) {
	file, _ := newConversationFile("test_data/v2.yml")
	_, err := file.V1Config()

	assert.NotNil(t, err)
}
