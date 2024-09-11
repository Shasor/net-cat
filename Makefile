all: build run

build:
	@echo "Build in progress..."
	@go build -o TCPChat ./cmd/main.go

run:
	@./TCPChat