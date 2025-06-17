.PHONY: test test-coverage coverage-preview clean

# Default target
all: test

# Run all tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...

# Preview coverage in browser
coverage-preview: test-coverage
	go tool cover -html=coverage.out

# Clean up generated files
clean:
	rm -f coverage.out 