BINARY_NAME := ghosted
.DEFAULT_GOAL := build
GO := go

build:
	$(GO) build -o ./bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

clean:
	rm ./bin/$(BINARY_NAME)

