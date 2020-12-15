-include ./env
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


test:
	$(GO_TEST) -v ./server/... ./crypto/...

build-server:
	$(GO_BUILD) $(SERVER_FILES)

build-client:
	cd $(CLIENT_SRC) && $(CLIENT_BUILD)

run-server:
	$(GO_RUN) $(SERVER_FILES)

run-client:
	cd $(CLIENT_SRC) && $(CLIENT_RUN)

run-all:
	$(GO_RUN) $(SERVER_FILES) & cd $(CLIENT_SRC) && $(CLIENT_RUN) && fg
