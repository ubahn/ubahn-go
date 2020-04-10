package core

type IInputService interface {
	Resolve(input interface{}) IResolvedInput
}

type IResolvedInput interface {
	Name() string
	Context() interface{}
}

type ResolvedInput struct {
}
