.DEFAULT_GOAL := build
.PHONY: fmt vet build clean
clean:
	go clean -i
fmt: clean
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
