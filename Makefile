PROJECT_NAME := "product-search-challenge"
PKG := "gitlab.com/a.vandam/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
CODE_FOLDER := src
MOCK_DB := "products-db"

.PHONY: all  clean test test-v help coverage coverhtml lint build-product-search

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test-all: ## Run unittests
	@go test -short ${PKG_LIST}

test-all-v:
	go test ${PKG_LIST} -v

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME) && docker-compose rm

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start-test-env:
	@docker-compose kill && make start-mock-db && docker-compose up --remove-orphans --build

start-mock-db:
	@cd $(MOCK_DB) && make database-reset && cd ../

stop-mock-db:
	@cd $(MOCK_DB) && make database-down
