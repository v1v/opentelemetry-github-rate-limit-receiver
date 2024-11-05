SHELL := bash
MAKEFLAGS += --no-print-directory
GITHUB_TOKEN ?= $(shell gh auth token)

#######################
## Tools
#######################
export PATH := $(CURDIR)/bin:$(PATH)
OCB ?= $(CURDIR)/bin/builder
METADATA ?= $(CURDIR)/bin/mdatagen

## @help:install-ngrok:Install ngrok.
.PHONY: install-ngrok
install-ngrok:
ifeq ($(OS),Darwin)
	brew install ngrok/ngrok/ngrok
else
	$(error "Please install ngrok manually")
endif

## @help:install-ocb:Install ocb.
.PHONY: install-ocb
install-ocb:
	GOBIN=$(CURDIR)/bin go install go.opentelemetry.io/collector/cmd/builder@v0.102.0

## @help:install-mdatagen:Install mdatagen.
.PHONY: install-mdatagen
install-mdatagen:
	GOBIN=$(CURDIR)/bin go install github.com/open-telemetry/opentelemetry-collector-contrib/cmd/mdatagen@latest

## MAKE GOALS

.PHONY: githubratelimit/metadata.go
githubratelimit/metadata.go: ## Build the metadata
githubratelimit/metadata.go: githubratelimit/metadata.yaml install-mdatagen
	cd githubratelimit && $(METADATA) metadata.yaml

.PHONY: build
build: githubratelimit/metadata.go ## Build the binary
	@$(OCB) --config builder-config.yml

.PHONY: test - Run tests for githubratelimit
test: 
	cd githubratelimit && go test -v ./...

.PHONY: run
run: ## Run the binary
	@GITHUB_TOKEN=$(GITHUB_TOKEN) \
	./bin/otelcol-custom --config config.yml
