FROM golang:1.13 AS builder

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
COPY go.* /build/

WORKDIR /build
RUN go mod download
RUN go mod verify

# Copy the source code into the container.
COPY pkg pkg
COPY apps apps

RUN go build \
    -mod=readonly \
    -ldflags "-X github.com/mmrath/gobase/golang/pkg/version.Version=$VERSION -X github.com/mmrath/gobase/golang/pkg/version.GitCommit=$GIT_COMMIT -X 'github.com/mmrath/gobase/golang/pkg/version.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'" \
    -a -o /go/bin/"$APP" ./apps/"$APP"

FROM scratch
COPY --from=builder /go/bin/"$APP" /app/"$APP"
ENTRYPOINT ["/app/${APP}"]

