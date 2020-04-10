package core

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// IConversationFile defines a file that contains conversation configuration.
type IConversationFile interface {
	Parse(out interface{}) error
}

// ConversationFile contains conversation configuration data, which can be parsed in
// to a structured object.
type ConversationFile struct {
	Version int
	Empty   bool
	Data    []byte
}

type conversationFileHeader struct {
	Version int `yaml:"version"`
}

var nullConversationFile = &ConversationFile{Empty: true}

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

	return &ConversationFile{Version: header.Version, Data: data}, nil
}

// Parse attempts to convert file content to the specified structure.
func (file *ConversationFile) Parse(out interface{}) error {
	return parseFileContent(file.Data, out)
}

func parseFileContent(content []byte, out interface{}) error {
	return yaml.Unmarshal(content, out)
}
