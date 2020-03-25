package core

// IOutput describes an output that can be sent to those in a conversation,
// who provided input (e.g. a chatbot sends an output to a user).
type IOutput interface {
	// Name returns output's name.
	Name() string

	// Send sends this output and returns a result.
	Send() (interface{}, error)
}

// BlankOutputName is a predefined constant name of a blank output that acts as a null-object
// in cases when output can't be returned.
const BlankOutputName string = "ubahn-blank-output"

// NotFoundOutputName is a predefined constant name of an output that indicates that no appropriate
// output was found.
const NotFoundOutputName string = "ubahn-not-found"

// NextOutputName is a predefined constant name which acts as a keyword, identifying the next
// output in the sequence.
const NextOutputName string = "next"

type nullOutput struct {
	name string
}

// NewNullOutput creates a null object that implements IOutput.
func NewNullOutput(name string) IOutput {
	return &nullOutput{name: name}
}

// BlankOutput is a null object of output.
var BlankOutput = NewNullOutput(BlankOutputName)

// NotFoundOutput is a null object of not found output.
var NotFoundOutput = NewNullOutput(NotFoundOutputName)

func (out *nullOutput) Name() string {
	return out.name
}

func (out *nullOutput) Send() (interface{}, error) {
	return nil, nil
}
