# Makefile for ollama-go

.PHONY: build test lint clean examples

# Build the library
build:
	go build ./...

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linting
lint:
	go vet ./...
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Build examples
examples:
	go build -o bin/chat ./examples/chat
	go build -o bin/completion ./examples/completion
	go build -o bin/embeddings ./examples/embeddings
	go build -o bin/models ./examples/models
	go build -o bin/tools ./examples/tools

# Run a specific example
run-chat: build
	go run ./examples/chat

run-completion: build
	go run ./examples/completion

run-embeddings: build
	go run ./examples/embeddings

run-models: build
	go run ./examples/models

run-tools: build
	go run ./examples/tools

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install golangci-lint if not installed
release:
    @echo "Creating new release $(VERSION)"
    @git tag -a v$(VERSION) -m "Release v$(VERSION)"
    @git push origin v$(VERSION)

# Help command
help:
	@echo "Available targets:"
	@echo "  build          - Build the library"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  lint           - Run linting"
	@echo "  release        - Create a new release"
	@echo "  examples       - Build all examples"
	@echo "  run-chat       - Run chat example"
	@echo "  run-completion - Run completion example"
	@echo "  run-embeddings - Run embeddings example"
	@echo "  run-models     - Run models example"
	@echo "  run-tools      - Run tools example"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
