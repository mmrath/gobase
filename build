#!/usr/bin/env bash

echo "building db_migration"
(cd db_migration; go build)



echo "building external api"
(cd external/api/app; wire; go build)

