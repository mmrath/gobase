FROM golang:1.13 AS buildenv

LABEL maintainer="Murali Mohan Rath <murali@mmrath.com>"

ARG APP
ARG VERSION
ARG GIT_COMMIT


ENV APP=${APP}
ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /build

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go/go.* /build/

WORKDIR /build
RUN go mod download
RUN go mod verify

# Copy the source code into the container.
COPY go/pkg pkg
COPY go/apps apps
COPY scripts scripts

RUN scripts/go-build.sh
