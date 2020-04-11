package core

// IExternalInputResolver describes a resolver that takes external input context and
// tries to resolve it to a known input.
type IExternalInputResolver interface {
	Resolve(context IExternalInputContext) IResolvedInput
}

type fakeResolveFunc func(msg string) IResolvedInput

// FakeExternalInputResolver is a fake implementation of external input resolver for testing purposes.
type FakeExternalInputResolver struct {
	resolveFunc fakeResolveFunc
}

// NewFakeExternalInputResolver creates a new instance of the fake input resolver.
func NewFakeExternalInputResolver(resolveFunc fakeResolveFunc) IExternalInputResolver {
	return &FakeExternalInputResolver{resolveFunc: resolveFunc}
}

// Resolve applies the internal resolve function and assumes external context data is a string.
func (resolver *FakeExternalInputResolver) Resolve(context IExternalInputContext) IResolvedInput {
	return resolver.resolveFunc(context.Data().(string))
}
