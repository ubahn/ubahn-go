package ubahn

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
	assert.Equal(t, 2, file.Version)
}

func Test_NewConversationFile_InvalidFile(t *testing.T) {
	file, err := newTestConversationFile("weird", "invalid.yml")

	assert.True(t, file.Empty())
	assert.NotNil(t, err)
}

func Test_NewConversationFile_WhenFileDoesNotExist(t *testing.T) {
	file, err := newTestConversationFile(defaultTestConversationName, "")

	assert.True(t, file.Empty())
	assert.NotNil(t, err)
}

func Test_NewConversationFile_Version(t *testing.T) {
	file, _ := newDefaultTestConversationFile()

	assert.Equal(t, 2, file.Version)
}

func Test_Parse(t *testing.T) {
	file, _ := newDefaultTestConversationFile()
	var config flowConfig

	err := file.Parse(&config)

	assert.Nil(t, err)
	_, hasKey := config.Outputs[0]["default-welcome"]
	assert.True(t, hasKey)
}

func Test_FileName(t *testing.T) {
	file, _ := newDefaultTestConversationFile()

	assert.Equal(t, defaultTestConversationFileName, file.FileName())
}

func Test_FilePath(t *testing.T) {
	file, _ := newDefaultTestConversationFile()

	assert.Equal(t,
		fmt.Sprintf("./test_data/%s/flows/%s", defaultTestConversationName, defaultTestConversationFileName),
		file.FilePath())
}

const defaultTestConversationName = "weather"
const defaultTestConversationFileName = "default.yml"

func newDefaultTestConversationFile() (*ConversationFile, error) {
	return newTestConversationFile(defaultTestConversationName, defaultTestConversationFileName)
}
