PROJECT_NAME := "ubahn-go"
PKG := "github.com/ubahn/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

lint:
	${GOPATH}/bin/golint -set_exit_status ${PKG_LIST}

test:
	@go test -short ${PKG_LIST} -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

race:
	@go test -race -short ${PKG_LIST}

msan:
	@go test -msan -short ${PKG_LIST}