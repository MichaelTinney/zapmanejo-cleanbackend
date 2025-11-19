.PHONY: help tidy build test lint docker-build clean run

# Default target
help:
	@echo "Available targets:"
	@echo "  make tidy         - Run go mod tidy"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make lint         - Run golangci-lint"
	@echo "  make docker-build - Build Docker image"
	@echo "  make run          - Run the application locally"
	@echo "  make clean        - Clean build artifacts"

# Run go mod tidy
tidy:
	go mod tidy

# Build the application
build:
	go build -v -o bin/zapmanejo-backend .

# Run tests
test:
	go test -v ./...

# Run linter
lint:
	golangci-lint run --timeout=5m

# Build Docker image
docker-build:
	docker build -t zapmanejo-backend:latest .

# Run the application locally
run:
	go run .

# Clean build artifacts
clean:
	rm -rf bin/
	go clean
