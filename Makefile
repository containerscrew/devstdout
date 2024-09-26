SHELL:=/bin/sh
.PHONY: all

help: ## this help
	@awk 'BEGIN {FS = ":.*?## ";  printf "Usage:\n  make \033[36m<target> \033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?## / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

mtoc: ## Create table of contents with mtoc
	mtoc

golangci-lint: ## Lint Golang code (brew install golangci-lint)
	golangci-lint run --fix

pre-commit: ## Run pre-commit
	pre-commit run -a

test: ## Run tests
	go test ./...

generate-changelog: ## Generate changelog using git cliff
	git cliff --output CHANGELOG.md
