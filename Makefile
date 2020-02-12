APPNAME:=$(shell basename $(shell go list))
COMMIT:=$(shell git describe --tags --always --dirty)
DATE:=$(shell date +%FT%T%z)
RELEASE?=0

ifeq ($(origin VERSION), undefined)
# check if there are any existing `git tag` values
ifeq ($(shell git tag),)
# no tags found - default to initial tag `v0.0.0`
VERSION := $(shell echo "v0.0.0-$$(git rev-list HEAD --count)-$$(git describe --dirty --always)" | sed 's/-/./2' | sed 's/-/./2')
else
# use tags
VERSION := $(shell git describe --dirty --always --tags | sed 's/-/./2' | sed 's/-/./2' )
endif
endif
export VERSION


GO_PROJ_BASE := $(shell pwd)
GOPATH ?= $(GO_PROJ_BASE)/.build
export GOPATH

GO_PROJ_BIN := $(GO_PROJ_BASE)/bin
GOBIN := $(GO_PROJ_BIN)

export PATH := $(GOBIN):$(PATH)

#Helps us keep project specific tool binaries

GO := GOBIN=$(GOBIN) go

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

GOLANGCILINT_VERSION ?= 1.22.2
GOLANGCILINT := $(GOBIN)/golangci-lint

GO_LDFLAGS+=-X github.com/mmrath/gobase/pkg/version.Version=$(VERSION)
GO_LDFLAGS+=-X github.com/mmrath/gobase/pkg/version.GitCommit=$(COMMIT)
GO_LDFLAGS+=-X github.com/mmrath/gobase/pkg/version.BuildTime=$(DATE)

ifeq ($(RELEASE), 1)
	# Strip debug information from the binary
	GO_LDFLAGS+=-s -w
endif
GO_LDFLAGS:=-ldflags="$(GO_LDFLAGS)"



.PHONY: clean
clean:
	@echo "  >  Cleaning build cache"
	@$(GO) clean

build-all: tools generate format lint build

build: generate
	@echo "Building version $(VERSION)"
	@$(GO) build $(GO_LDFLAGS) -o $(GO_PROJ_BIN) ./apps/...

.PHONY: lint
lint: $(GOLANGCILINT)
	@echo "Running lint"
	@$(GOLANGCILINT) run

.PHONY: format
format:
	@echo "Running format"
	@gofmt -s -w ./apps ./pkg ./*.go

.PHONY: generate
generate: tools
	@echo "Running go generate"
	@$(GO) generate ./...

.PHONY: test
test:
	@$(GO) test -v ./...

tools: $(GOLANGCILINT)
	@echo Installing tools
	@$(GO) install github.com/cespare/reflex
	@$(GO) install github.com/go-bindata/go-bindata/...

$(GOLANGCILINT):
	@curl -fsSL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v$(GOLANGCILINT_VERSION)

.PHONY: go
go:
	@$(GO) $(CMD)

.PHONY: docker-clean
docker-clean:
	docker rm -v $$(docker ps -aq -f 'status=exited')
	docker rmi $$(docker images -aq -f 'dangling=true')
	docker volume rm $$(docker volume ls -q -f 'dangling=true')


