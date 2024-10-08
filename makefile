# Makefile for Go JSON config reader

# Variables
BINARY_NAME=IBIS
GO=go
GOBUILD=$(GO) build
GORUN=$(GO) run
GOINSTALL=$(GO) install
GOTEST=$(GO) test -v
GOMOD=$(GO) mod tidy
GOFMT=$(GO) fmt

# Main source file
SOURCE=ibis.go

# Default target
all: build

# Build the application
build:
	$(GOFMT)
	$(GOMOD)
	$(GOBUILD) -o $(BINARY_NAME)
	migrate -source file://migrations -database ${IBIS_DATABASE_URL} up

# Run the application
run: build
	./$(BINARY_NAME)

# Run the tests verbosely
test: build
	$(GOTEST)

# Clean up binary
clean:
	rm -f $(BINARY_NAME)

db:
	fly postgres connect -a ibisdb

ssh:
	fly ssh console

# Phony targets
.PHONY: all build run test clean db ssh