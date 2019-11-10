package core

// IOutputFactory describes a factory that creates output objects.
type IOutputFactory interface {
	Create(outputName string) IOutput
}