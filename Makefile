# Makefile to build go-sdk-template
include ./make/config.mk

all: build unittest lint tidy

travis-ci: install build alltest lint tidy

build:
	go build ./...

unittest:
	go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./... | grep -v samples`

alltest: export PACT_TEST := true
alltest:
	go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./... | grep -v samples` -v -tags=integration
	ls -al pacts

lint:
	golangci-lint run

tidy:
	go mod tidy

install:
	@if [ ! -d pact/bin ]; then\
		echo "--- Installing Pact CLI dependencies";\
		curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash;\
	fi

publish: install
	@echo "--- 📝 Publishing Pacts"
	go run sqlv2/pact/publish.go
	@echo
	@echo "Pact contract publishing complete!"
