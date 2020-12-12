-include ./env
GOBIN=go
BUILD_SRC=./server/*.go
TIMEOUT=20

test:
	$(GOBIN) test -v ./server/... ./crypto/...

build:
	$(GOBIN) build $(BUILD_SRC)

run:
	$(GOBIN) run $(BUILD_SRC)


