SHELL := /usr/bin/env bash -o pipefail

PROJECT := wellknowntypes
GO_MODULE := github.com/bufbuild/wellknowntypes
BUF_VERSION := 1.0.0-rc11

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
CACHE_BIN := $(CACHE)/bin
CACHE_VERSIONS := $(CACHE)/versions
CACHE_GO := $(CACHE)/go
CACHE_ENV := $(CACHE)/env
CACHE_GOBIN := $(CACHE)/gobin
CACHE_GOCACHE := $(CACHE)/gocache

export GO111MODULE := on
ifdef GOPRIVATE
export GOPRIVATE := $(GOPRIVATE),$(GO_MODULE)
else
export GOPRIVATE := $(GO_MODULE)
endif
export GOPATH := $(abspath $(CACHE_GO))
export GOBIN := $(abspath $(CACHE_GOBIN))
export GOCACHE := $(abspath $(CACHE_GOCACHE))
export GOMODCACHE := $(GOPATH)/pkg/mod

EXTRAPATH := $(abspath $(GOBIN)):$(abspath $(CACHE_BIN))
export PATH := $(EXTRAPATH):$(PATH)

BUF := $(CACHE_VERSIONS)/buf/$(BUF_VERSION)
$(BUF):
	@rm -f $(CACHE_BIN)/buf
	@mkdir -p $(CACHE_BIN)
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"
	@rm -rf $(dir $(BUF))
	@mkdir -p $(dir $(BUF))
	@touch $(BUF)

.DEFAULT_GOAL := all

.PHONY: all
all:
	$(MAKE) populate
	$(MAKE) breaking

.PHONY: direnv
direnv:
	@mkdir -p $(CACHE_ENV)
	@rm -f $(CACHE_ENV)/env.sh
	@echo 'export CACHE="$(abspath $(CACHE))"' >> $(CACHE_ENV)/env.sh
	@echo 'export GO111MODULE="$(GO111MODULE)"' >> $(CACHE_ENV)/env.sh
	@echo 'export GOPRIVATE="$(GOPRIVATE)"' >> $(CACHE_ENV)/env.sh
	@echo 'export GOPATH="$(GOPATH)"' >> $(CACHE_ENV)/env.sh
	@echo 'export GOBIN="$(GOBIN)"' >> $(CACHE_ENV)/env.sh
	@echo 'export GOCACHE="$(GOCACHE)"' >> $(CACHE_ENV)/env.sh
	@echo 'export GOMODCACHE="$(GOPATH)/pkg/mod"' >> $(CACHE_ENV)/env.sh
	@echo 'export PATH="$(EXTRAPATH):$${PATH}"' >> $(CACHE_ENV)/env.sh
	@echo $(CACHE_ENV)/env.sh

.PHONY: install
install:
	@go install ./internal/cmd/...

.PHONY: populate
populate: install
	bash internal/script/populate.bash

.PHONY: breaking
breaking: install $(BUF)
	bash internal/script/breaking.bash

.PHONY: printversions
printversions: install
	@ls | grep ^v3 | sort-semver-tags

.PHONY: clean
clean:
	git clean -xdf
	rm -rf $(CACHE_BASE)
