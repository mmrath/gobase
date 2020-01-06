PROJ=gobase
ORG_PATH=github.com/mmrath
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)
GIT_COMMIT := $(shell git rev-list -1 HEAD)


$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOAPP=$(PWD)/bin


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

version: ## Show version
	@echo $(VERSION) \(git commit: $(GIT_COMMIT)\)

start:
	docker-compose -f docker-compose.local.yml up --build

# DOCKER TASKS
docker-build: ## [DOCKER] Build given container. Example: `make docker-build APP=clipo`
	docker build -f go/Dockerfile --no-cache --build-arg APP=$(APP) --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) -t go-build:local .

docker-run: ## [DOCKER] Run container on given port. Example: `make docker-run APP=user PORT=3000`
	docker run -i -t "$(APP):local" --rm -p=$(PORT):$(PORT) --name="$(APP)" $(APP)

docker-stop: ## [DOCKER] Stop docker container. Example: `make docker-stop APP=user`
	docker stop $(APP)

docker-rm: docker-stop ## [DOCKER] Stop and then remove docker container. Example: `make docker-rm APP=user`
	docker rm $(APP)

docker-publish: docker-repo-login docker-publish-latest docker-publish-version ## [DOCKER] Docker publish. Example: `make docker-publish APP=user REGISTRY=https://your-registry.com`

docker-publish-latest: docker-tag-latest
	@echo 'publish latest to $(REGISTRY)'
	docker push $(REGISTRY)/$(APP):latest

docker-publish-version: docker-tag-version
	@echo 'publish $(VERSION) to $(REGISTRY)'
	docker push $(REGISTRY)/$(APP):$(VERSION)

docker-tag: docker-tag-latest docker-tag-version ## [DOCKER] Tag current container. Example: `make docker-tag APP=user REGISTRY=https://your-registry.com`

docker-tag-latest:
	@echo 'create tag latest'
	docker tag $(APP) $(REGISTRY)/$(APP):latest
	docker tag $(APP) $(REGISTRY)/$(APP):latest

docker-tag-version:
	@echo 'create tag $(VERSION)'
	docker tag $(APP) $(REGISTRY)/$(APP):$(VERSION)

docker-release: docker-build docker-publish ## [DOCKER] Docker release - build, tag and push the container. Example: `make docker-release APP=user REGISTRY=https://your-registry.com`


docker-repo-login: ## [HELPER] login to docker repo
	@echo "run script/cmd to login to docker repo"

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
         -subj "/C=AU/ST=NSW/L=Sydney/O=Sample SSL Certificate/CN=localhost"
	@openssl req \
         -newkey rsa:2048 -nodes -keyout dist/key_pair/sso_private.key \
         -x509 -days 365 -out dist/key_pair/sso_public.crt \
         -subj "/C=AU/ST=NSW/L=Sydney/O=Sample SSL Certificate/CN=localhost"

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







