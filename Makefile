PROJECT_NAME := "product-search-challenge"
PKG := "gitlab.com/a.vandam/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
CODE_FOLDER := src
MOCK_DB := "products-db"
ENV_FILE := .env


.PHONY: all  clean test test-v help coverage coverhtml lint build-product-search

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test-all: ## Run unittests
	@go test -short ${PKG_LIST}

test-all-v:
	go test ${PKG_LIST} -v

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME) && docker-compose rm -f

start-svc:
	@make clean && docker-compose up --build --

start-test-env:
	@docker-compose kill && make start-mock-db && docker-compose up --build 	

start-mock-db:
	@cd $(MOCK_DB) && make database-reset && docker network connect product-search-challenge_default mongodb-local  && cd ../

stop-mock-db:
	@cd $(MOCK_DB) && make database-down
