#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

    CREATE ROLE clipo WITH LOGIN ENCRYPTED PASSWORD '${CLIPO_DB_PASSWORD}';
    CREATE ROLE oppo WITH LOGIN ENCRYPTED PASSWORD '${OPPO_DB_PASSWORD}';
    CREATE ROLE db_migration WITH LOGIN ENCRYPTED PASSWORD '${MIGRATION_DB_PASSWORD}';

    REVOKE ALL ON SCHEMA public FROM PUBLIC;
    GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} TO db_migration;
    GRANT ALL PRIVILEGES ON SCHEMA public TO db_migration;

    ALTER ROLE db_migration IN DATABASE ${POSTGRES_DB} SET search_path TO public;
    ALTER ROLE clipo IN DATABASE ${POSTGRES_DB} SET search_path TO public;
    ALTER ROLE oppo IN DATABASE ${POSTGRES_DB} SET search_path TO public;

EOSQL