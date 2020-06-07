package ubahn

/// OUTPUTS

// BlankOutputName is a predefined constant name of a blank output that acts as a null-object
// in cases when output can't be returned.
// It also can be used in cases when no output is needed.
const BlankOutputName string = ":blank"

// NotFoundOutputName is a predefined constant name of an output that indicates that no appropriate
// output was found.
const NotFoundOutputName string = ":not-found"

// NextOutputName is a predefined constant name which acts as a keyword, identifying the next
// output in the sequence.
const NextOutputName string = ":next"
