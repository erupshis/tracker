
.PHONY: build-container docker_up docker_down
build-container:
	docker build -t tracker -f build/Dockerfile .

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

