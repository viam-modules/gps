BIN_OUTPUT_PATH = bin
TOOL_BIN = bin/gotools/$(shell uname -s)-$(shell uname -m)
COMMON_LDFLAGS = -s -w #-X 'go.viam.com/rdk/config.Version=${TAG_VERSION}' -X 'go.viam.com/rdk/config.GitRevision=${GIT_REVISION}' -X 'go.viam.com/rdk/config.DateCompiled=${DATE_COMPILED}'
UNAME_S ?= $(shell uname -s)

ifeq ($(shell command -v dpkg >/dev/null && dpkg --print-architecture),armhf)
GOFLAGS += -tags=no_tflite
endif

module: build
	rm -f $(BIN_OUTPUT_PATH)/module.tar.gz
	tar czf $(BIN_OUTPUT_PATH)/module.tar.gz $(BIN_OUTPUT_PATH)/gps meta.json

build: build-go

build-go:
	rm -f $(BIN_OUTPUT_PATH)/gps
	go build -tags no_cgo,osusergo,netgo -ldflags="-extldflags=-static $(COMMON_LDFLAGS)"  -o $(BIN_OUTPUT_PATH)/gps main.go

# Install each tool by its own module version. One `go install` cannot take
# packages from different modules, and installing them inside this module pulls
# linter libraries that golangci-lint v1.62.0 cannot build against.
# golangci-lint and the combined vet tool must be compiled with this module's
# Go toolchain. The other tools keep the toolchain selected by their own go.mod.
tool-install:
	GOTOOLCHAIN=go$$(go list -f '{{.GoVersion}}' -m) GOBIN=`pwd`/$(TOOL_BIN) \
		go install github.com/edaniels/golinters/cmd/combined@v0.0.5-0.20220906153528-641155550742
	GOTOOLCHAIN=go$$(go list -f '{{.GoVersion}}' -m) GOBIN=`pwd`/$(TOOL_BIN) \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0
	GOBIN=`pwd`/$(TOOL_BIN) go install github.com/AlekSi/gocov-xml@v1.0.0
	GOBIN=`pwd`/$(TOOL_BIN) go install github.com/axw/gocov/gocov@v1.1.0
	GOBIN=`pwd`/$(TOOL_BIN) go install gotest.tools/gotestsum@v1.12.2
	GOBIN=`pwd`/$(TOOL_BIN) go install github.com/rhysd/actionlint/cmd/actionlint@v1.7.8

lint: lint-go
	PATH=$(TOOL_BIN) actionlint

lint-go: tool-install
	go mod tidy
	export pkgs="`go list -f '{{.Dir}}' ./... | grep -v /proto/`" && echo "$$pkgs" | xargs go vet -vettool=$(TOOL_BIN)/combined
	GOGC=50 $(TOOL_BIN)/golangci-lint run -v --fix --config=./etc/golangci.yaml

test: test-go

test-go: tool-install
	go test -race ./...

clean-all:
	git clean -fxd

license-check:
	license_finder