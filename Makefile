# Makefile for Blockchain project

# Run the application
run:
	go run main.go

# Run all tests
test:
	go test ./blockchain/...

# Build the application
build:
	go build -o blockchain-app main.go

# Clean up build files
clean:
	rm -f blockchain-app