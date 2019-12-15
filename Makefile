PROJ=gobase
ORG_PATH=github.com/mmrath
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

build: bin/uaa-server bin/uaa-client-example bin/db_migration

bin/uaa-server: uaa/uaa-server
	go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/uaa/uaa-server

bin/uaa-client-example:
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/uaa/uaa-client-example

bin/db_migration:
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/db_migration



test:
	@go test -v ./...

testrace:
	@go test -v --race ./...

vet:
	@go vet ./...

fmt:
	@./scripts/gofmt ./...

lint: bin/golint
	@./bin/golint -set_exit_status $(shell go list ./...)



.PHONY: docker-image
docker-image:
	@sudo docker build -t $(DOCKER_IMAGE) .

.PHONY: proto
proto: bin/protoc bin/protoc-gen-go
	@./bin/protoc --go_out=plugins=grpc:. --plugin=protoc-gen-go=./bin/protoc-gen-go api/*.proto
	@./bin/protoc --go_out=. --plugin=protoc-gen-go=./bin/protoc-gen-go server/internal/*.proto

.PHONY: verify-proto
verify-proto: proto
	@./scripts/git-diff

bin/protoc: scripts/get-protoc
	@./scripts/get-protoc bin/protoc

bin/protoc-gen-go:
	@go install -v $(REPO_PATH)/vendor/github.com/golang/protobuf/protoc-gen-go

bin/golint:
	@go install -v $(REPO_PATH)/vendor/golang.org/x/lint/golint

clean:
	@rm -rf bin/

testall: testrace vet fmt lint

FORCE:

.PHONY: test testrace vet fmt lint testall

generate_certs:
	@mkdir -p dist/ssl_certs
	@mkdir -p dist/key_pair
	@openssl req \
         -newkey rsa:2048 -nodes -keyout dist/ssl_certs/ssl_private.key \
         -x509 -days 365 -out dist/ssl_certs/ssl_public.crt \
         -subj "/C=CA/ST=British Columbia/L=Vancouver/O=Sample SSL Certificate/CN=localhost"
	@openssl req \
         -newkey rsa:2048 -nodes -keyout dist/key_pair/sso_private.key \
         -x509 -days 365 -out dist/key_pair/sso_public.crt \
         -subj "/C=CA/ST=British Columbia/L=Vancouver/O=Sample SSL Certificate/CN=localhost"

gen_go_repo:
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_repositories_new.bzl%go_repositories
	@rm go_repositories.bzl
	@mv go_repositories_new.bzl go_repositories.bzl
	@bazel run //:gazelle

uaa:
	@bazel build apps/uaa

uaa-admin:
	@bazel build apps/uaa-admin

db-migration:
	@bazel build apps/db-migration

apps: uaa uaa-admin db-migration