.PHONY: all build mocks test

commands  = web
webpack   = public/assets/docs.js

assets   = $(wildcard assets/*)
binaries = $(addprefix $(GOPATH)/bin/, $(commands))
compiler = go
mode     = development
sources  = $(shell find . -name '*.go')

ifeq ($(PACKAGE),true)
	compiler=go run -mod=vendor vendor/github.com/gobuffalo/packr/packr/main.go
	mode=production
endif

all: build

build: $(binaries) $(webpack)

mocks:
	make -C models mocks
	make -C pkg/storage mocks

test:
	env TEST=true go test -coverpkg ./... -covermode atomic -coverprofile coverage.txt ./...

$(binaries): $(GOPATH)/bin/%: $(sources)
	$(compiler) install ./cmd/$*

$(GOPATH)/bin/web: $(webpack)

$(webpack): $(assets)
	node webpack/node_modules/webpack/bin/webpack.js --config webpack/webpack.config.js --mode $(mode)
