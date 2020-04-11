package ubahn

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	core "github.com/ubahn/ubahn-go/core"
)

const welcomeInputName = "i-welcome"
const yesInputName = "i-yes"
const numInputName = "i-num"

var dummyTextResolver core.IExternalInputResolver = core.NewFakeExternalInputResolver(
	func(msg string) core.IResolvedInput {
		switch msg {
		case "hi":
			return core.NewResolvedInput(welcomeInputName, msg)
		case "yes":
			return core.NewResolvedInput(yesInputName, msg)
		default:
			return core.UnresolvedInput
		}
	})

var dummyNumResolver core.IExternalInputResolver = core.NewFakeExternalInputResolver(
	func(msg string) core.IResolvedInput {
		if ok, _ := regexp.MatchString(`\d+`, msg); ok {
			return core.NewResolvedInput(numInputName, msg)
		}

		return core.UnresolvedInput
	})

func Test_Resolve_WhenEmpty(t *testing.T) {
	assert.False(t, optimisticResolve([]core.IExternalInputResolver{}, "hi").IsResolved())
}

func Test_Resolve(t *testing.T) {
	resolvers := []core.IExternalInputResolver{dummyTextResolver, dummyNumResolver}

	assert.Equal(t, welcomeInputName, optimisticResolve(resolvers, "hi").Name())
	assert.Equal(t, yesInputName, optimisticResolve(resolvers, "yes").Name())
	assert.Equal(t, numInputName, optimisticResolve(resolvers, "123").Name())

	unknown := optimisticResolve(resolvers, "blah blah")
	assert.Equal(t, "", unknown.Name())
	assert.False(t, unknown.IsResolved())
}

func optimisticResolve(resolvers []core.IExternalInputResolver, msg string) core.IResolvedInput {
	resolver := NewOptimisticInputResolver(resolvers)
	return resolver.Resolve(core.NewFakeExternalInputContext(msg))
}
