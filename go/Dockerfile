FROM golang:1.13 AS builder

LABEL maintainer="Murali Mohan Rath <murali@mmrath.com>"

ARG VERSION
ARG GIT_COMMIT

ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /build

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /build/
COPY *.go /build/
COPY Makefile /build/

WORKDIR /build


# Copy the source code into the container.
COPY pkg pkg
COPY apps apps

RUN make build-all

FROM alpine:latest AS db-migration
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/bin/db-migration .
CMD ["./db-migration"]

FROM alpine:latest AS clipo
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/bin/clipo .
CMD ["./clipo"]

FROM alpine:latest AS admin
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/bin/admin .
CMD ["./admin"]

