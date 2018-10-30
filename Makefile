NATIVE_BIN := build/$(shell go env GOOS)/$(shell go env GOARCH)/tfplan2json

.PHONY: all
all: $(NATIVE_BIN)

.PHONY: clean
clean:
	$(RM) -r ./build

.PHONY: distclean
distclean: clean
	$(RM) vendor/.deps-fetched

.PHONY: deps
deps: vendor/.deps-fetched

vendor/.deps-fetched:
	gvt rebuild
	touch $@

$(NATIVE_BIN): $(wildcard *.go)
	go build -o $@ .
