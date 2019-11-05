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
build: client master

client:
	$(GOBUILD) -o $(GOBIN)/client ./cmd/client/ 

master:
	$(GOBUILD) -o $(GOBIN)/master ./cmd/master

clean:
	rm -rf ./bin

fmt:
	$(GOFMT) ./cmd/client ./cmd/master

