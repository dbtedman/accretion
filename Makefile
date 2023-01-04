.DEFAULT_GOAL := all

.PHONY: all
all: install lint type_check build test sast

.PHONY: pre_commit
pre_commit: lint type_check test_fast sast

.PHONY: install
install: install_go install_pnpm

.PHONY: install_go
install_go:
	@go mod vendor

.PHONY: install_pnpm
install_pnpm:
	@pnpm install

.PHONY: lint
lint:
	@pnpm run lint && gofmt -l ./cmd ./internal && golangci-lint run

.PHONY: type_check
type_check:
	@pnpm run type-check

.PHONY: format
format:
	@pnpm run format && gofmt -w ./cmd ./internal && golangci-lint run --fix

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

.PHONY: sast
sast: sast_snyk sast_osv

.PHONY: sast_snyk
sast_snyk:
	snyk test --all-projects --detection-depth=1

.PHONY: sast_osv
sast_osv:
	osv-scanner ./
