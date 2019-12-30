FROM scratch
COPY --from=builder /go/bin/db-migration /app/db-migration
ENTRYPOINT ["/app/db-migration"]

