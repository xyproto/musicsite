.PHONY: build clean run

# Variables
SERVER_DIR=./cmd/server
FRONTEND_DIR=./frontend
WASM_FILE=$(FRONTEND_DIR)/main.wasm
GO_EXECUTABLE=go
WASM_EXEC_JS=$(shell go env GOROOT)/misc/wasm/wasm_exec.js

# Default target
all: build

# Build the project
build: wasm

# Build WebAssembly binary
wasm:
	@echo "Building WebAssembly module..."
	GOOS=js GOARCH=wasm $(GO_EXECUTABLE) build -o $(WASM_FILE) $(FRONTEND_DIR)/main.go
	cp $(WASM_EXEC_JS) $(FRONTEND_DIR)

# Clean up the build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(WASM_FILE)
	rm -f $(FRONTEND_DIR)/wasm_exec.js

# Run the server
run:
	@echo "Running server on http://localhost:8080..."
	$(GO_EXECUTABLE) run cmd/server/main.go
