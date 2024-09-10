all: build run

build:
	@echo "Build in progress..."
	@go build -o TCPChat ./cmd/main.go

run:
	@echo "Run:"
	@go run ./cmd/main.go