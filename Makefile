#!/usr/bin/make -f

########################################################
#                       Makefile                       #
########################################################

# Default target
.DEFAULT_GOAL := build
all: build

MODULES := $(shell find . -type f -name 'go.mod' -exec dirname {} \;)

########################################################
#                         Setup                        #
########################################################

# Generate versioning information
TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")
VERSION := $(TAG:v%=%)
ifneq ($(COMMIT), $(TAG_COMMIT))
    VERSION := $(VERSION)-next-$(COMMIT)-$(DATE)
endif
ifneq ($(shell git status --porcelain),)
    VERSION := $(VERSION)-dirty
endif

########################################################
#                       Building                       #
########################################################

# Build all services
build:
	@echo Building go-pyth-client
	@go build -o bin/go-pyth-client feeds.go

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
	@go list -f '{{.Dir}}/...' -m | xargs golangci-lint run  --timeout=10m --concurrency 8 -v 

golangci-fix:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | xargs golangci-lint run  --timeout=10m --fix --concurrency 8 -v 

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

test-go:
	@$(MAKE) install-ginkgo
	@ginkgo -r --vv --randomize-all --fail-on-pending -trace ./...

test:
	@$(MAKE) test-go

test-go-cover: 
	@export GOCOVERDIR=$(shell pwd)
	@$(MAKE) install-ginkgo
	@ginkgo -r --vv --randomize-all --fail-on-pending --trace \
		--cover --coverprofile "coverage-go.txt" --covermode atomic ./...

test-cover:
	@$(MAKE) test-go-cover

########################################################
#                        Dependency                    #
########################################################

repo-rinse: |
	git submodule foreach --recursive git clean -xfd
	git submodule foreach --recursive git reset --hard
	git submodule update --init --recursive

tidy: |
	@for module in $(MODULES); do \
		echo "Running go mod tidy in $$module"; \
		(cd $$module && go mod tidy) || exit 1; \
	done