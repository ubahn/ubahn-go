package ubahn

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

type OutputFactory struct {
}

func NewOutputFactory(services map[string]IOutputService) IOutputFactory {
	return &OutputFactory{}
}

func (factory *OutputFactory) Create(outputName string) IOutput {
	return nil
}
