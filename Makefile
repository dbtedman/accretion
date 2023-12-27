.DEFAULT_GOAL := all

.PHONY: all
all: install test

.PHONY: install
install:
	@go mod tidy && go mod vendor

.PHONY: test
test:
	@CGO_ENABLED=0 go test ./event/...
