FILES		  ?= $(shell find . -type f -name '*.go')

.PHONY: fmt
fmt: ## format the go source files
	go fmt ./...
	goimports -local "github.com/mibk/dupl" -w $(FILES)

.PHONY: lint
lint: ## run go lint on the source files
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Run 'make tools'"; \
		exit 1; \
	fi; \
	golangci-lint run ./... --max-issues-per-linter 10 --max-same-issues 50

.PHONY: tools
tools: ## fetch and install all required tools
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.28.3

