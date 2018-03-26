NATIVE_BIN := build/$(shell go env GOOS)/$(shell go env GOARCH)/tfplan2json

all: build

build: $(NATIVE_BIN)

$(NATIVE_BIN):
	go build -o $@ .
