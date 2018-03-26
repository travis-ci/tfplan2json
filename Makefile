NATIVE_BIN := build/$(shell go env GOOS)/$(shell go env GOARCH)/tfplan2json

.PHONY: all
all: $(NATIVE_BIN)

.PHONY: clean
clean:
	rm -rf ./build

$(NATIVE_BIN): $(wildcard *.go)
	go build -o $@ .
