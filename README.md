# Ubahn for Golang

Golang implementation of the conversation structuring framework Ubahn.
See the [whitepaper](https://github.com/ubahn/whitepaper) for generic information on Ubahn.

[![Build Status](https://travis-ci.org/ubahn/ubahn-go.svg?branch=master)](https://travis-ci.org/ubahn/ubahn-go)
[![Build status](https://ci.appveyor.com/api/projects/status/xo85r9pinevo74f1?svg=true)](https://ci.appveyor.com/project/slavikdev/ubahn-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/7c6ac2dd052e2c817a90/maintainability)](https://codeclimate.com/github/ubahn/ubahn-go/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/ubahn/ubahn-go)](https://goreportcard.com/report/github.com/ubahn/ubahn-go)
[![GoDoc](https://godoc.org/github.com/ubahn/ubahn-go?status.svg)](https://godoc.org/github.com/ubahn/ubahn-go)

## Installation and usage

The import path for the package is *gopkg.in/ubahn/ubahn-go.v1*.

To install it, run:

    go get gopkg.in/ubahn/ubahn-go.v1

## Example

Conversation definition in `weather-report.yml`:

```yaml
version: 1

sequence:
  - welcome
  - weather-report

triggers:
  - i-user-welcome

fallback: clarification

outputs:
  welcome:
    expectedInputs:
      i-yes: next
      i-no: bye
    fallback: welcome-clarification
```

Business logic:

```go
// inputFactory should be specific to your app and produce instances of ubahn.IInput.
// Of course you may produce them in a different way.
// Ubahn doesn’t provide input factory interface because it doesn’t want to assume input format.
inputFactory = NewInputFactory()

// outputFactory should implement ubahn.IOutputFactory
// interface and be specific to your app.
outputFactory = NewOutputFactory()

// This can be initialized once. Conversation object is stateless.
conversation := ubahn.NewConversation("weather-report.yml")

// For the sake of simplicity in this example we define user input as an array
// and just loop through it. Of course in your app you’ll receive it properly from
// real users.
userInputs := []string{"Hi", "Yes"}

// We store previous output, because conversation object is stateless.
// ubahn.BlankOutputName indicates the beginning of the conversation.
prevOutputName := ubahn.BlankOutputName

for i := 0; i < len(userInputs); i++ {
  input := inputFactory.Create(userInputs[i])
  outputName := conversation.Continue(prevOutputName, input)
  output := outputFactory.Create(outputName)
  output.Send()
  prevOutputName = outputName
}
```

Here’s a checklist of what you should implement on your app’s side in order to use Ubahn:

* Definition of conversations in YAML format.
* Something that produces instances of ubahn.IInput based on user input (like a factory).
* Specific output factory that implements ubahn.IOutputFactory.
* Specific outputs that implement ubahn.IOutput.
