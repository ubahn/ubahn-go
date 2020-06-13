package ubahn

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// IConversationFile defines a file that contains conversation configuration.
type IConversationFile interface {
	Parse(out interface{}) error
	FileName() string
	FilePath() string
	Empty() bool
}

// ConversationFile contains conversation configuration data, which can be parsed in
// to a structured object.
type ConversationFile struct {
	Version  int
	empty    bool
	Data     []byte
	path     string
	fileName string
}

type conversationFileHeader struct {
	Version int `yaml:"version"`
}

var nullConversationFile = &ConversationFile{empty: true}

// NewConversationFile creates a configuration file from a given full path.
// It attempts to read version from the file header and also read file content.
func NewConversationFile(path string) (*ConversationFile, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nullConversationFile, err
	}
	data := []byte(fileContent)
	header := &conversationFileHeader{}
	err = parseFileContent(data, header)
	if err != nil {
		return nullConversationFile, err
	}

	return &ConversationFile{
		Version:  header.Version,
		Data:     data,
		path:     path,
		fileName: filepath.Base(path)}, nil
}

// Parse attempts to convert file content to the specified structure.
func (file *ConversationFile) Parse(out interface{}) error {
	return parseFileContent(file.Data, out)
}

// FileName returns the name of the conversation file without the full path.
func (file *ConversationFile) FileName() string {
	return file.fileName
}

// FilePath returns the full path of the conversation file including the file name.
func (file *ConversationFile) FilePath() string {
	return file.path
}

// Empty returns true if the file is empty or not initialized, otherwise returns false.
func (file *ConversationFile) Empty() bool {
	return file.empty
}

func parseFileContent(content []byte, out interface{}) error {
	return yaml.Unmarshal(content, out)
}
