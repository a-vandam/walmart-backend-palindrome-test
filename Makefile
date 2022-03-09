PROJECT_NAME := "product-search-challenge"
PKG := "gitlab.com/a.vandam/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
CODE_FOLDER := src

.PHONY: all  clean test test-v help coverage coverhtml lint build-product-search

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-v:
	go test ${PKG_LIST} -v

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test-env:
	@docker-compose kill && docker-compose up --remove-orphans --build
