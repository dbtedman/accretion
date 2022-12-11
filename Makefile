.DEFAULT_GOAL := all

.PHONY: all
all: install lint build

.PHONY: pre_commit
pre_commit: lint test_fast

.PHONY: install
install: install_go
	@pnpm install

.PHONY: install_go
install_go:
	@go mod vendor

.PHONY: lint
lint:
	@pnpm run lint && gofmt -l ./cmd ./internal

.PHONY: format
format:
	@pnpm run format && gofmt -w ./cmd ./internal

.PHONY: test
test:
	@pnpm run test && go test -race -cover -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: test_fast
test_fast:
	@pnpm run test:fast && CGO_ENABLED=0 go test ./...

.PHONY: build
build: build_ts build_go

.PHONY: build_ts
build_ts:
	@pnpm run build

.PHONY: build_go
build_go: clean_go
	@CGO_ENABLED=0 go build -mod vendor -o ./bin/accretion ./cmd/accretion

.PHONY: clean_go
clean_go:
	@rm -rf ./bin

.PHONY: upgrade
upgrade:
	@pnpm dlx npm-check-updates -u && pnpm upgrade && go get -u all && go mod tidy

local:
	@docker compose down --volumes --rmi local \
		&& docker compose up --detach \
		&& echo "visit http://localhost:3000"
