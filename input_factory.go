package ubahn

import (
	core "github.com/ubahn/ubahn-go/core"
)

type InputFactory struct {
}

func NewInputFactory(services map[string]core.IInputService) *InputFactory {
	return &InputFactory{}
}
