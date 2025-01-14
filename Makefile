.PHONY: go-mod-tidy gofumpt goimports golangci-lint gotestsum buf-generate buf-lint buf-format buf-breaking \
        go-rewrite-flow go-test-flow go-ci-flow all-go-tasks \
        buf-test-flow buf-ci-flow all-buf-tasks all-tasks

# Go-related tasks
go-mod-tidy:
	go mod tidy

gofumpt:
	gofumpt -extra -w .

goimports:
	find . -name "*.go" -not -path "./gen/*" -exec goimports -w -local github.com/cresplanex/bloader {} +

golangci-lint:
	golangci-lint run

gotestsum:
	gotestsum --format=short-verbose

# Buf-related tasks
buf-generate:
	buf generate

buf-lint:
	buf lint

buf-format:
	buf format -w

buf-breaking:
	buf breaking proto --against '.git#subdir=proto'

# Combined Go flows
go-rewrite-flow: gofumpt goimports
	@echo "Go rewrite flow completed."

go-test-flow: golangci-lint gotestsum
	@echo "Go test flow completed."

go-ci-flow: go-rewrite-flow go-test-flow
	@echo "Go CI flow completed."

all-go-tasks: go-mod-tidy go-ci-flow
	@echo "All Go tasks completed."

# Combined Buf flows
buf-test-flow: buf-lint buf-breaking
	@echo "Buf test flow completed."

buf-ci-flow: buf-format buf-test-flow
	@echo "Buf CI flow completed."

all-buf-tasks: buf-generate buf-ci-flow
	@echo "All Buf tasks completed."

# All tasks
all-tasks: all-buf-tasks all-go-tasks
	@echo "All tasks completed."
