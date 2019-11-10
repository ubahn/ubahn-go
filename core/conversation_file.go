package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type IConversationFile interface {
	Parse(out interface{}) error
}

type ConversationFile struct {
	Version int
	Empty   bool
	Data    []byte
}

type conversationFileHeader struct {
	Version int `yaml:"version"`
}

var nullConversationFile = &ConversationFile{Empty: true}

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

func (file *ConversationFile) Parse(out interface{}) error {
	return parseFileContent(file.Data, out)
}

func parseFileContent(content []byte, out interface{}) error {
	return yaml.Unmarshal(content, out)
}
