PROJECT_NAME := "product-search-challenge"
PKG := "gitlab.com/a.vandam/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
CODE_FOLDER := src
MOCK_DB := "products-db"
#docker variables
MOCK_DB_CONTAINER_NAME := "mongodb-local"
TEST_NETWORK := "project-network"
ENV_FILE := .env


.PHONY: all  clean test test-v help coverage coverhtml lint build-product-search

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME) && docker-compose rm -f && make stop-mock-db && make stop-svc && docker network rm ${TEST_NETWORK}

start-svc: 
	@make clean && docker-compose up --build --

stop-svc:
	@docker-compose kill $(PROJECT_NAME)

start-test-env:
	@docker-compose kill && make start-mock-db && docker-compose up --build 	

start-mock-db:
	@cd $(MOCK_DB) && make database-reset && cd ../ && \
	docker network create ${TEST_NETWORK} && \
	docker network connect ${TEST_NETWORK} ${MOCK_DB_CONTAINER_NAME}

stop-mock-db:
	@cd $(MOCK_DB) && make database-down

svc-docker-build:
	@docker build -t "${PROJECT_NAME}" . 

svc-docker-run:
	@docker run --env-file ${ENV_FILE} --network=${TEST_NETWORK} --rm -it ${PROJECT_NAME} 