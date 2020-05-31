package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConversationFile(t *testing.T) {
	file, err := newDefaultTestConversationFile()

	assert.False(t, file.Empty())
	assert.Nil(t, err)
	assert.NotEmpty(t, file.Data)
	assert.Equal(t, 1, file.Version)
}

func Test_NewConversationFile_InvalidFile(t *testing.T) {
	file, err := newTestConversationFile("v1", "invalid.yml")

	assert.True(t, file.Empty())
	assert.NotNil(t, err)
}

func Test_NewConversationFile_WhenFileDoesNotExist(t *testing.T) {
	file, err := newTestConversationFile("v1", "")

	assert.True(t, file.Empty())
	assert.NotNil(t, err)
}

func Test_NewConversationFile_Version(t *testing.T) {
	file, _ := newTestConversationFile("v1", "v1.yml")

	assert.Equal(t, 1, file.Version)

	file, _ = newTestConversationFile("v1", "v2.yml")

	assert.Equal(t, 2, file.Version)
}

func Test_Parse(t *testing.T) {
	file, _ := newDefaultTestConversationFile()
	var hash map[string]interface{}
	err := file.Parse(&hash)

	assert.Nil(t, err)

	sequence := hash["sequence"].([]interface{})
	assert.Equal(t, 2, len(sequence))
	assert.Equal(t, "welcome", sequence[0].(string))
	assert.Equal(t, "weather-report", sequence[1].(string))
}

func Test_FileName(t *testing.T) {
	file, _ := newDefaultTestConversationFile()

	assert.Equal(t, defaultConversationFileName, file.FileName())
}

func Test_FilePath(t *testing.T) {
	file, _ := newDefaultTestConversationFile()

	assert.Equal(t, "../test_data/v1/"+defaultConversationFileName, file.FilePath())
}

func newTestConversationFile(version, fileName string) (*ConversationFile, error) {
	return NewConversationFile(
		fmt.Sprintf("../test_data/%s/%s", version, fileName))
}

const defaultConversationFileName = "weather.yml"

func newDefaultTestConversationFile() (*ConversationFile, error) {
	return newTestConversationFile("v1", defaultConversationFileName)
}
