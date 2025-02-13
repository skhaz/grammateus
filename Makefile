.PHONY: help fmt lint test update vet
.SILENT:

SHELL := bash -eou pipefail

ifeq ($(shell command -v docker-compose;),)
		COMPOSE := docker compose
else
		COMPOSE := docker-compose
endif

help:
	awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt:
	npx prettier --write --ignore-unknown --ignore-path .gitignore --print-width=120 pkg/handler/public/index.html
	go fmt ./...

lint: vet ## Run lint
	golangci-lint run ./...

test: ## Run tests
	go test ./...

update: ## Upgrade dependencies
	go get -u -t -v ./...
	go mod tidy

vet: ## Run vet
	go vet ./...
