#!/usr/bin/env bash

if [ -z "$1" ]; then
    echo "ERROR Must provide application name(clipo, admin) as argument"
    exit 1
fi

if [ "$1" != "clipo" ] && [ "$1" != "admin" ]; then
  echo "ERROR Argument 1 must be in (clipo, admin)"
  exit 1
fi


if [ -z "$VERSION" ]; then
    VERSION=$(git rev-parse HEAD)
fi

if [ -z "$GIT_COMMIT" ]; then
    GIT_COMMIT=$VERSION
fi

APP=$1

echo "Building application ${APP}"

BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S')
GO_LD_FLAGS="-X github.com/mmrath/gobase/go/pkg/version.Version=$VERSION -X github.com/mmrath/gobase/go/pkg/version.GitCommit=$GIT_COMMIT -X 'github.com/mmrath/gobase/go/pkg/version.BuildTime=$BUILD_TIME'"

echo "Downloading go modules"
go mod download
echo "Downloaded go modules"
go install -ldflags "${GO_LD_FLAGS}" "./apps/${APP}"

retval=$?
if [ $retval -ne 0 ]; then
    echo "FAILED: Building application ${APP} "
else
    echo "SUCCESS: Building application ${APP}"
fi
