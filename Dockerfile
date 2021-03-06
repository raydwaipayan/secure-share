# Build Stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . ./

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o bin/server server/*.go

# Production stage
FROM scratch
WORKDIR /app

COPY --from=builder /app/bin/server /app/server
ENTRYPOINT ["/app/server"]