.PHONY: help
help: ## Show the help
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the project
	go run cmd/httpechoes/main.go

snapshot: ## Snapshot release with GoReleaser
	goreleaser build --snapshot --clean

.PHONY: lint
lint: ## Lint Golang codebase
	golangci-lint run

.PHONY: lint-todo
lint-todo: ## Find TODOs in Golang codebase
	golangci-lint run --no-config --disable-all --enable godox

.PHONY: fmt
fmt: ## Format Golang codebase and "optimize" the dependencies
	golangci-lint fmt
	go mod tidy
