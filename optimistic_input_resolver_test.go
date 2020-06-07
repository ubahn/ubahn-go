package ubahn

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

const welcomeInputName = "i-welcome"
const yesInputName = "i-yes"
const numInputName = "i-num"

var dummyTextResolver IExternalInputResolver = NewFakeExternalInputResolver(
	func(msg string) IResolvedInput {
		switch msg {
		case "hi":
			return NewResolvedInput(welcomeInputName, msg)
		case "yes":
			return NewResolvedInput(yesInputName, msg)
		default:
			return UnresolvedInput
		}
	})

var dummyNumResolver IExternalInputResolver = NewFakeExternalInputResolver(
	func(msg string) IResolvedInput {
		if ok, _ := regexp.MatchString(`\d+`, msg); ok {
			return NewResolvedInput(numInputName, msg)
		}

		return UnresolvedInput
	})

func Test_Resolve_WhenEmpty(t *testing.T) {
	assert.False(t, optimisticResolve([]IExternalInputResolver{}, "hi").IsResolved())
}

func Test_Resolve(t *testing.T) {
	resolvers := []IExternalInputResolver{dummyTextResolver, dummyNumResolver}

	assert.Equal(t, welcomeInputName, optimisticResolve(resolvers, "hi").Name())
	assert.Equal(t, yesInputName, optimisticResolve(resolvers, "yes").Name())
	assert.Equal(t, numInputName, optimisticResolve(resolvers, "123").Name())

	unknown := optimisticResolve(resolvers, "blah blah")
	assert.Equal(t, "", unknown.Name())
	assert.False(t, unknown.IsResolved())
}

func optimisticResolve(resolvers []IExternalInputResolver, msg string) IResolvedInput {
	resolver := NewOptimisticInputResolver(resolvers)
	return resolver.Resolve(NewFakeExternalInputContext(msg))
}
