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
	go build -tags no_cgo,osusergo,netgo -ldflags="-extldflags=-static $(COMMON_LDFLAGS)" -o $(BIN_OUTPUT_PATH) main.go

build-win:
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -tags no_cgo,osusergo,netgo -ldflags="-extldflags=-static $(COMMON_LDFLAGS)" -o bin/gps.exe .
	cat meta.json | jq '.entrypoint |= "bin/gps.exe"' > meta-win.json
	mv meta.json meta.backup.json
	mv meta-win.json meta.json
	tar czf bin/module.tar.gz bin/gps.exe meta.json

tool-install:
	GOBIN=`pwd`/$(TOOL_BIN) go install \
		github.com/edaniels/golinters/cmd/combined \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/AlekSi/gocov-xml \
		github.com/axw/gocov/gocov \
		gotest.tools/gotestsum \
		github.com/rhysd/actionlint/cmd/actionlint

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
