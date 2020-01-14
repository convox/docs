.PHONY: all build mocks package test

commands  = web
webpack   = public/assets/docs.js

assets   = $(wildcard assets/*)
binaries = $(addprefix $(GOPATH)/bin/, $(commands))
mode     = development
sources  = $(shell find . -name '*.go')

all: build

build: $(binaries) $(webpack)

mocks:
	make -C models mocks
	make -C pkg/storage mocks

package:
	go run -mod=vendor vendor/github.com/gobuffalo/packr/packr/main.go

test:
	env TEST=true go test -coverpkg ./... -covermode atomic -coverprofile coverage.txt ./...

$(binaries): $(GOPATH)/bin/%: $(sources)
	go install -mod=vendor --ldflags="-s -w" ./cmd/$*

$(GOPATH)/bin/web: $(webpack)

$(webpack): $(assets)
	node webpack/node_modules/webpack/bin/webpack.js --config webpack/webpack.config.js --mode $(mode)
