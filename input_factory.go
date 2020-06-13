package ubahn

// IInputFactory describes a factory that creates input objects, based on external context.
type IInputFactory interface {
	Create(context IExternalInputContext) IInput
}
