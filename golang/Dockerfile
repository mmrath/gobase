FROM golang:1.13 AS builder

LABEL maintainer="Murali Mohan Rath <murali@mmrath.com>"

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /work

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /work/
COPY *.go /work/
COPY Makefile /work/

WORKDIR /work


# Copy the source code into the container.
COPY pkg pkg
COPY apps apps

RUN make build

FROM alpine:latest AS db-migration
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /work/bin/db-migration .
CMD ["./db-migration"]

FROM alpine:latest AS clipo
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /work/bin/clipo .
CMD ["./clipo"]

FROM alpine:latest AS admin
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /work/bin/oppo .
CMD ["./oppo"]