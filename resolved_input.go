package ubahn

// IResolvedInput describes a result of external input resolver, which can be later used to
// create an input object.
type IResolvedInput interface {
	Name() string
	IsResolved() bool
	Data() interface{}
}

// ResolvedInput is the default implementation of the IResolvedInput interface.
type ResolvedInput struct {
	name string
	data interface{}
}

// NewResolvedInput creates a new instance of the default resolved input type.
func NewResolvedInput(name string, data interface{}) IResolvedInput {
	return &ResolvedInput{name: name, data: data}
}

// Name returns resolved input name.
func (input *ResolvedInput) Name() string {
	return input.name
}

// IsResolved returns true if the input was resolved and false otherwise.
// In the default implementation we consider an input resolved, if it has a name, which means a resolver
// was able to find a name for it.
func (input *ResolvedInput) IsResolved() bool {
	return len(input.Name()) > 0
}

// Data returns resolved input data. Can be some useful information, such as original user message.
func (input *ResolvedInput) Data() interface{} {
	return input.data
}

// UnresolvedInput is a predefined result which is returned when resolver can’t resolve input.
var UnresolvedInput = NewResolvedInput("", nil)
