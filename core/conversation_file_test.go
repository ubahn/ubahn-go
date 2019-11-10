package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConversationFile(t *testing.T) {
	file, err := NewConversationFile("../test_data/weather.yml")

	assert.False(t, file.Empty)
	assert.Nil(t, err)
	assert.NotEmpty(t, file.Data)
	assert.Equal(t, 1, file.Version)
}

func TestNewConversationFile_InvalidFile(t *testing.T) {
	file, err := NewConversationFile("../test_data/invalid.yml")

	assert.True(t, file.Empty)
	assert.NotNil(t, err)
}

func TestNewConversationFile_WhenFileDoesNotExist(t *testing.T) {
	file, err := NewConversationFile("")

	assert.True(t, file.Empty)
	assert.NotNil(t, err)
}

func TestNewConversationFile_Version(t *testing.T) {
	file, _ := NewConversationFile("../test_data/v1.yml")

	assert.Equal(t, 1, file.Version)

	file, _ = NewConversationFile("../test_data/v2.yml")

	assert.Equal(t, 2, file.Version)
}

func TestParse(t *testing.T) {
	file, _ := NewConversationFile("../test_data/weather.yml")
	var hash map[string]interface{}
	err := file.Parse(&hash)

	assert.Nil(t, err)

	sequence := hash["sequence"].([]interface{})
	assert.Equal(t, 2, len(sequence))
	assert.Equal(t, "welcome", sequence[0].(string))
	assert.Equal(t, "weather-report", sequence[1].(string))
}
