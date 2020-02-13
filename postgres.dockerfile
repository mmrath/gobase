FROM postgres:11-alpine
EXPOSE 5432
COPY ./db/*.sql /docker-entrypoint-initdb.d/