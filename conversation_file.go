package ubahn

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conversationConfigHeader struct {
	Version int `yaml:"version"`
}

type configV1Output struct {
	ExpectedInputs map[string]string `yaml:"expectedInputs"`
	Fallback       string            `yaml:"fallback"`
	empty          bool
}

type conversationConfigV1 struct {
	Sequence []string                  `yaml:"sequence"`
	Triggers []string                  `yaml:"triggers"`
	Fallback string                    `yaml:"fallback"`
	Outputs  map[string]configV1Output `yaml:"outputs"`
}

type conversationFile struct {
	version int
	empty   bool
	data    []byte
}

func newConversationFile(path string) (*conversationFile, error) {
	nullObject := &conversationFile{empty: true}
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nullObject, err
	}
	data := []byte(fileContent)
	header := &conversationConfigHeader{}
	err = parse(data, header)
	if err != nil {
		return nullObject, err
	}

	return &conversationFile{version: header.Version, data: data}, nil
}

func (file *conversationFile) V1Config() (*conversationConfigV1, error) {
	if file.version != 1 {
		return &conversationConfigV1{}, errors.New("file version is not 1")
	}

	config := &conversationConfigV1{}
	err := parse(file.data, &config)
	return config, err
}

func parse(content []byte, out interface{}) error {
	return yaml.Unmarshal(content, out)
}
