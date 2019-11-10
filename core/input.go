package core

// IInput defines an input which comes from the outer conversation party (e.g. chatbot user).
type IInput interface {
	// Name returns input name.
	Name() string
}
