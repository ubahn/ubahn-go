# Ubahn for Golang

Golang implementation of the dialog management framework Ubahn.
See the [whitepaper](https://github.com/ubahn/whitepaper) for generic information on Ubahn.

[![Build Status](https://travis-ci.org/ubahn/ubahn-go.svg?branch=master)](https://travis-ci.org/ubahn/ubahn-go)
[![Build status](https://ci.appveyor.com/api/projects/status/xo85r9pinevo74f1?svg=true)](https://ci.appveyor.com/project/slavikdev/ubahn-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/7c6ac2dd052e2c817a90/maintainability)](https://codeclimate.com/github/ubahn/ubahn-go/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/ubahn/ubahn-go)](https://goreportcard.com/report/github.com/ubahn/ubahn-go)
[![GoDoc](https://godoc.org/github.com/ubahn/ubahn-go?status.svg)](https://pkg.go.dev/github.com/ubahn/ubahn-go)

## Overview

Ubahn currently uses rule-based approach to dialog management, allowing developers to specify conversation flows
in `yaml` files. Chatbot application that uses Ubahn may apply machine learning to understand user input.
Moreover developers can use ML to generate outputs. Ubahn’s purpose is to match input with output, based on
predefined configuration. It’s especially useful when a chatbot application has to follow certain steps to guide users.

## Installation

To install the package run:

    go get github.com/ubahn/ubahn-go

## Next steps

- [ ] Create examples and documentation
- [ ] Create contribution guidelines
- [ ] Move towards ML vs rule based approach
