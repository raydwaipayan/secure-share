#!make
include .env
export $(shell sed 's/=.*//' .env)

TIMEOUT=20

GOBIN=go
GO_BUILD=$(GOBIN) build
GO_RUN=$(GOBIN) run
GO_TEST=$(GOBIN) test -timeout $(TIMEOUT)

SERVER_SRC=./server
SERVER_FILES=$(SERVER_SRC)/*.go

CLIENT_SRC=./client
CLIENT_BUILD=yarn build
CLIENT_RUN=yarn serve

.PHONY : test build-server build-client run-server run-client run-all copy-env deploy

test:
	$(GO_TEST) -v ./server/... ./crypto/...

build-server:
	$(GO_BUILD) -o bin/server $(SERVER_FILES)

build-client:
	cd $(CLIENT_SRC) && $(CLIENT_BUILD)

clean:
	rm -rf bin

run-server:
	$(GO_RUN) $(SERVER_FILES)

copy-env:
	cp .env client/.env

run-client: copy-env
	cd $(CLIENT_SRC) && $(CLIENT_RUN)

run-all: copy-env
	$(GO_RUN) $(SERVER_FILES) & cd $(CLIENT_SRC) && $(CLIENT_RUN) && fg

deploy: copy-env
	docker-compose up -d