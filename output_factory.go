package ubahn

import (
	core "github.com/ubahn/ubahn-go/core"
)

type OutputFactory struct {
}

func NewOutputFactory(services map[string]core.IOutputService) core.IOutputFactory {
	return &OutputFactory{}
}

func (factory *OutputFactory) Create(outputName string) core.IOutput {
	return nil
}
