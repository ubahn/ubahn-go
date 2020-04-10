package core

// IOutputFactory describes a factory that creates output objects.
type IOutputFactory interface {
	Create(outputName string) IOutput
}

// NullOutputFactory creates null outputs.
type NullOutputFactory struct {
}

// NewNullOutputFactory creates an instance of NullOutputFactory.
func NewNullOutputFactory() IOutputFactory {
	return &NullOutputFactory{}
}

// Create creates a null output.
func (factory *NullOutputFactory) Create(outputName string) IOutput {
	return NewNullOutput(outputName)
}
