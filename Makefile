all: build

build:
	@go build -o mqtt-analyzer

run: build
	@./mqtt-analyzer