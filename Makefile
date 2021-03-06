OUT := bin/sentimentapi
PKG := github.com/brozeph/sentimentapi
BLD := $(shell date +%FT%T%z)
VER := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

# to discover variable paths for linking: go tool nm <path_to_binary> | grep <variable_name>

all: run

server:
	@go build -i -o ${OUT} -ldflags="\
		-X ${PKG}/internal/resources.Version=${VER} \
		-X ${PKG}/internal/resources.Build=${BLD} \
		-X ${PKG}/internal/resources.Package=${PKG}" ${PKG}/cmd

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	@go build -i -o ${OUT}-v${VER} -tags netgo -ldflags="-extldflags \"-static\" -w -s \
		-X ${PKG}/internal/resources.Version=${VER} \
		-X ${PKG}/internal/resources.Build=${BLD} \
		-X ${PKG}/internal/resources.Package=${PKG}" ${PKG}/cmd

run: server
	./${OUT}

clean:
	@rm -r bin
	@mkdir bin

.PHONY: run server static vet lint