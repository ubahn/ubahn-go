package ubahn

// OptimisticInputResolver is an external input resolver that uses other specific resolvers
// to resolve input.
type OptimisticInputResolver struct {
	resolvers []IExternalInputResolver
}

// NewOptimisticInputResolver creates a new instance of optimistic input resolver.
func NewOptimisticInputResolver(resolvers []IExternalInputResolver) IExternalInputResolver {
	return &OptimisticInputResolver{resolvers: resolvers}
}

// Resolve goes through the internal collection of resolvers tries all of them until it gets a resolved input.
// If all resolvers failed, it returns unresolved input.
func (resolver *OptimisticInputResolver) Resolve(context IExternalInputContext) IResolvedInput {
	for i := 0; i < len(resolver.resolvers); i++ {
		input := resolver.resolvers[i].Resolve(context)
		if input.IsResolved() {
			return input
		}
	}
	return UnresolvedInput
}
