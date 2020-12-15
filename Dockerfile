FROM golang

WORKDIR /app
COPY . ./

RUN go build -o bin/server server/*.go
ENTRYPOINT ./bin/server

# Document that the service listens on port 8080.
EXPOSE 8080