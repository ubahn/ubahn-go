package ubahn

type IOutputService interface {
	Resolve(output interface{}) IResolvedOutput
}

type IResolvedOutput interface {
	Name() string
	Context() interface{}
}

type ResolvedOutput struct {
}
