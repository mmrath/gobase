#!/usr/bin/env bash

if [ -z "$1" ]; then
    echo "ERROR Must provide application name(clipo, admin) as argument"
    exit 1
fi

if [ "$1" != "clipo" ] && [ "$1" != "admin" ]; then
  echo "ERROR Argument 1 must be in (clipo, admin)"
  exit 1
fi


APP=$1

go install github.com/cespare/reflex

reflex -s -r "^(apps/${APP}|pkg)" -- sh -c "./build.sh ${APP} && ${GOPATH}/bin/${APP}"


