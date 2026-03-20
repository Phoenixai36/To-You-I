# T.U.I — To Unify Imagination
# Makefile — Build, run, test and install targets
# by Phoenixai36 · Dark Phoenix · Barcelona

BINARY   := tuicortex
CMD_PATH := ./cmd/tuicortex
BUILD_DIR := ./bin
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0-dev")
LDFLAGS  := -ldflags "-X main.version=$(VERSION)"

.PHONY: all build run clean test lint install tidy

## all: build the binary (default)
all: build

## build: compile the tuicortex binary
build:
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) $(CMD_PATH)
	@echo "Built $(BUILD_DIR)/$(BINARY) v$(VERSION)"

## run: build and run the TUI
run: build
	$(BUILD_DIR)/$(BINARY)

## run-dev: run directly without building (faster iteration)
run-dev:
	go run $(LDFLAGS) $(CMD_PATH)

## test: run all tests
test:
	go test ./... -v -race

## lint: run golangci-lint
lint:
	golangci-lint run ./...

## tidy: tidy go modules
tidy:
	go mod tidy

## install: install the binary to GOPATH/bin
install:
	go install $(LDFLAGS) $(CMD_PATH)
	@echo "Installed $(BINARY) to $(shell go env GOPATH)/bin"

## clean: remove build artifacts
clean:
	rm -rf $(BUILD_DIR)
	@echo "Cleaned $(BUILD_DIR)"

## help: show this help
help:
	@echo ""
	@echo "  T.U.I — To Unify Imagination"
	@echo "  Available targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'
	@echo ""
