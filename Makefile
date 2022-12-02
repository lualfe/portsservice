.DEFAULT_GOAL := help

.PHONY: help

IMAGE_TAG:="1.0"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-app: ## Executes the application. You need docker and docker compose installed.
	@docker compose up

stop-app: ## Stops the app. You need docker and docker compose installed.
	@docker compose down

build: ## Build app image. You need docker installed.
	@docker build -t portsservice:$(IMAGE_TAG) .

cleanup: ## Cleans up app containers. You need docker and docker compose installed.
	@docker compose rm

run-tests: ## Execute all tests
	@go test ./...

lint: ## Run linter. You need golangci-lint installed.
	@golangci-lint run