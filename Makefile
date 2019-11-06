 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin

# Build
build:
	$(GOBUILD) -mod vendor -o $(GOBIN)/tiger ./cmd/

clean:
	rm -rf ./bin

fmt:
	$(GOFMT) ./cmd

