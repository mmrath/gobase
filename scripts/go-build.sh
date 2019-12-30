#!/bin/sh

BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S')
GO_LD_FLAGS="-X github.com/mmrath/gobase/pkg/version.Version=$VERSION -X github.com/mmrath/gobase/pkg/version.GitCommit=$GIT_COMMIT -X github.com/mmrath/gobase/pkg/version.BuildTime='$BUILD_TIME'"

echo "$GO_LD_FLAGS"

go build -mod=readonly -ldflags "$GO_LD_FLAGS" -a -o ./dist/clipo ./apps/clipo