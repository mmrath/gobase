APPNAME:=$(shell basename $(shell go list))
VERSION?=snapshot
COMMIT:=$(shell git describe --tags --always --dirty)
DATE:=$(shell date +%FT%T%z)
RELEASE?=0



GOPATH?=$(shell go env GOPATH)
GO_LDFLAGS+=-X main.appName=$(APPNAME)
GO_LDFLAGS+=-X main.buildVersion=$(VERSION)
GO_LDFLAGS+=-X main.buildCommit=$(COMMIT)
GO_LDFLAGS+=-X main.buildDate=$(DATE)



generate_certs:
	@mkdir -p dist/jwt_certs
	@mkdir -p dist/key_pair
	@openssl req \
         -newkey rsa:2048 -nodes -keyout dist/jwt_certs/private.key \
         -x509 -days 365 -out dist/jwt_certs/public.crt \
         -subj "/C=AU/ST=NSW/L=Sydney/O=Sample SSL Certificate/CN=localhost"
	@openssl req \
         -newkey rsa:2048 -nodes -keyout dist/key_pair/sso_private.key \
         -x509 -days 365 -out dist/key_pair/sso_public.crt \
         -subj "/C=AU/ST=NSW/L=Sydney/O=Sample SSL Certificate/CN=localhost"











