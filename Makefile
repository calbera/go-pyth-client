#!/usr/bin/make -f

# Default target
.DEFAULT_GOAL := all
all: tidy bindings format lint test

MODULES := $(shell find . -type f -name 'go.mod' -exec dirname {} \;)

########################################################
#                       Building                       #
########################################################

# Generate solidity bindings for the Pyth EVM contracts
bindings:
	@echo "Reading the contracts build artifacts from bindings/abis/"
	cd bindings && go generate ./...

generate: bindings

###############################################################################
###                                Linting                                  ###
###############################################################################

format:
	@$(MAKE) golangci-fix

lint:
	@$(MAKE) golangci gosec

#################
# golangci-lint #
#################

golangci-install:
	@echo "--> Installing golangci-lint latest"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

golangci:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | xargs go run github.com/golangci/golangci-lint/cmd/golangci-lint run --config ./.golangci.yml --timeout=10m --concurrency 8 

golangci-fix:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | xargs go run github.com/golangci/golangci-lint/cmd/golangci-lint run --config ./.golangci.yml --timeout=10m --fix --concurrency 8 

#################
#     gosec     #
#################

gosec-install:
	@echo "--> Installing gosec latest"
	@go install github.com/securego/gosec/v2/cmd/gosec@latest

gosec:
	@$(MAKE) gosec-install
	@echo "--> Running gosec"
	@gosec -exclude-generated ./...


########################################################
#                        Testing                       #
########################################################

install-ginkgo:
	@echo "Installing ginkgo..."
	@go install github.com/onsi/ginkgo/v2/ginkgo@latest

test:
	@$(MAKE) install-ginkgo
	@ginkgo -r --vv --randomize-all --fail-on-pending -trace ./...

test-cover: 
	@export GOCOVERDIR=$(shell pwd)
	@$(MAKE) install-ginkgo
	@ginkgo -r --vv --randomize-all --fail-on-pending --trace \
		--cover --coverprofile "coverage-go.txt" --covermode atomic ./...

########################################################
#                        Dependency                    #
########################################################

tidy: |
	@for module in $(MODULES); do \
		echo "Running go mod tidy in $$module"; \
		(cd $$module && go mod tidy) || exit 1; \
	done
