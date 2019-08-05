OUT := bin/sentimentapi
PKG := github.com/brozeph/sentimentapi
BLD := $(shell date +%FT%T%z)
VER := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

# to discover variable paths for linking: go tool nm <path_to_binary> | grep <variable_name>

all: run

clean:
	@rm -r bin
	@mkdir bin

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

mongodb:
	./third_party/mongodb.sh

run: mongodb server
	./${OUT}

server:
	@go build -i -o ${OUT} -ldflags="\
		-X github.com/brozeph/sentimentapi/internal/resources.Version=${VER} \
		-X github.com/brozeph/sentimentapi/internal/resources.Build=${BLD} \
		-X github.com/brozeph/sentimentapi/internal/resources.Package=${PKG}" ${PKG}/cmd

static: vet lint
	@go build -i -o ${OUT}-v${VER} -tags netgo -ldflags="-extldflags \"-static\" -w -s \
	-X github.com/brozeph/sentimentapi/internal/resources.Version=${VER} \
	-X github.com/brozeph/sentimentapi/internal/resources.Build=${BLD} \
	-X github.com/brozeph/sentimentapi/internal/resources.Package=${PKG}" ${PKG}/cmd

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

.PHONY: run server static vet lint mongodb