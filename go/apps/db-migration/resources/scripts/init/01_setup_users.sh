#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE ROLE app_user WITH LOGIN ENCRYPTED PASSWORD '${APP_DB_PASSWORD}';
EOSQL