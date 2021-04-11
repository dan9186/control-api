SHELL = bash

APP := $(shell basename $(PWD) | tr '[:upper:]' '[:lower:]')
TTY := $(shell if [ -z $$CI ]; then echo "-t"; fi)

GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_SHORT_COMMIT_HASH := $(shell git rev-parse --short HEAD)

.PHONY: all
all: test

.PHONY: clean
clean: ## Cleans out all generated items
	-@rm -f output.txt
	-@rm -rf coverage
	-@rm -f coverage.txt
	-@docker-compose down

.PHONY: help
help:  ## Show This Help
	@for line in $$(cat Makefile | grep "##" | grep -v "grep" | sed  "s/:.*##/:/g" | sed "s/\ /!/g"); do verb=$$(echo $$line | cut -d ":" -f 1); desc=$$(echo $$line | cut -d ":" -f 2 | sed "s/!/\ /g"); printf "%-30s--%s\n" "$$verb" "$$desc"; done

.PHONY: proto
proto: ## Build the protobufs
	-@rm -rf api
	@protoc \
		-I ./proto \
		-I $(GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
		-I $(GOPATH)/src/github.com/googleapis/googleapis \
		--go_out="." \
		--go-grpc_out="." \
		--validate_out="lang=go:." \
		$$(find proto -type f -name '*.proto')
