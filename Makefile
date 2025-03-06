
include ./build/.env
export $(shell sed -n 's/^\([^#][^=]*\)=.*/\1/p' ./build/.env)

.PHONY: build-docker
build-docker:
	docker build \
		--build-arg GO_VERSION=${GO_VERSION} \
		--file ./build/Dockerfile \
		-t tracker \
		.

docker_down:
	docker compose -f ./build/docker-compose.yaml stop && docker compose -f ./build/docker-compose.yaml down

docker_up:
	docker compose -f ./build/docker-compose.yaml up

#////////////////_CHECK_CODE_/////////////////////
# Makefile for Go project with golangci-lint

GO_VERSION = 1.24
GOLANGCI_LINT_VERSION = v1.64.5
GOPATH_BIN := $(shell go env GOPATH)/bin

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_BIN) $(GOLANGCI_LINT_VERSION)

lint:
	$(GOPATH_BIN)/golangci-lint run --timeout=5m


#///////////////_MIGRATIONS_/////////////////////
MIGRATION_DIR=./migrations
DATABASE_URL=postgres://postgres:postgres@localhost:5432/trackerdb?sslmode=disable

# Check if MIGRATION_NAME is provided
check-migration-name:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required"; \
		exit 1; \
	fi

migration_create: check-migration-name
	goose -dir ./migrations create $(NAME) sql

migration_up:
	goose postgres $(DATABASE_URL) -dir $(MIGRATION_DIR) up

migration_down:
	goose postgres $(DATABASE_URL) -dir $(MIGRATION_DIR) down

migration_status:
	goose postgres $(DATABASE_URL) -dir $(MIGRATION_DIR) status