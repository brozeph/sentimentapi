OUT := dist/sentimentapi
PKG := github.com/brozeph/sentimentapi/cmd
BLD := $(shell date +%FT%T%z)
VER := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

# to discover variable paths for linking: go tool nm <path_to_binary> | grep <variable_name>

all: run

server:
	@go build -i -o ${OUT} -ldflags="\
		-X github.com/brozeph/sentimentapi/internal/resources.Version=${VER} \
		-X github.com/brozeph/sentimentapi/internal/resources.Build=${BLD} \
		-X github.com/brozeph/sentimentapi/internal/resources.Package=${PKG}" ${PKG}

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
	-X github.com/brozeph/sentimentapi/internal/resources.Version=${VER} \
	-X github.com/brozeph/sentimentapi/internal/resources.Build=${BLD} \
	-X github.com/brozeph/sentimentapi/internal/resources.Package=${PKG}" ${PKG}

run: server
	./${OUT}

clean:
	@rm -r dist
	@mkdir dist

.PHONY: run server static vet lint