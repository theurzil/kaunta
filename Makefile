.PHONY: build run dev clean install test

# Build the binary
build:
	go build -o kaunta .

# Run the application
run: build
	./kaunta

# Development mode with hot reload (requires air)
dev:
	air

# Clean build artifacts
clean:
	rm -f kaunta
	go clean

# Install dependencies
install:
	go mod download
	go mod tidy

# Run tests
test:
	go test -v ./...

# Install development tools
tools:
	go install github.com/air-verse/air@latest

# Database migrations (TODO)
migrate-up:
	@echo "TODO: Implement migrations"

migrate-down:
	@echo "TODO: Implement migrations"
