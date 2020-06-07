package ubahn

// IInput defines an input which comes from the outer conversation party (e.g. chatbot user).
type IInput interface {
	// Name returns input name.
	Name() string
}

type nullInput struct {
	name string
}

// NewNullInput creates a null object that implements IInput.
func NewNullInput(name string) IInput {
	return &nullInput{name: name}
}

func (in *nullInput) Name() string {
	return in.name
}
