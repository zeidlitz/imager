# Build stage
FROM docker.io/library/golang:1.21.5 as builder
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0
RUN go build ./cmd/kvdb/

# Run stage
FROM gcr.io/distroless/static-debian12:nonroot
EXPOSE 8080
COPY --from=builder /app/kvdb /app/kvdb
ENTRYPOINT ["/app/kvdb"]
