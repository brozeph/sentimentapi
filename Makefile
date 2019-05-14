OUT := sentimentapi
PKG := github.com/brozeph/sentimentapi
BLD := $(shell date +%FT%T%z)
VER := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

# to discover variable paths for linking: go tool nm <path_to_binary> | grep <variable_name>

all: run

server:
	@go build -i -v -o ${OUT} -ldflags="\
	-X github.com/brozeph/sentimentapi/resources.Version=${VER} \
	-X github.com/brozeph/sentimentapi/resources.Build=${BLD} \
	-X github.com/brozeph/sentimentapi/resources.Package=${PKG}" ${PKG}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	@go build -i -v -o ${OUT}-v${VER} -tags netgo -ldflags="-extldflags \"-static\" -w -s \
	-X github.com/brozeph/sentimentapi/resources.Version=${VER} \
	-X github.com/brozeph/sentimentapi/resources.Build=${BLD} \
	-X github.com/brozeph/sentimentapi/resources.Package=${PKG}" ${PKG}

run: server
	./${OUT}

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run server static vet lint