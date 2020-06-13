package ubahn

// IExternalInputContext describes context with data, received from external source.
// For example it may contain the original user message and information about it.
type IExternalInputContext interface {
	Data() interface{}
}

// FakeExternalInputContext is a fake implementation of external input context, for testing purposes.
type FakeExternalInputContext struct {
	data interface{}
}

// NewFakeExternalInputContext creates a new instance of the fake input context.
func NewFakeExternalInputContext(message string) IExternalInputContext {
	return &FakeExternalInputContext{data: message}
}

// Data simply returns incapsulated data object.
func (context *FakeExternalInputContext) Data() interface{} {
	return context.data
}
