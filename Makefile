NAMESPACE=briancain
NAME=gameboy-go
BINARY=${NAME}
# Builds
VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.1")
REF?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
GIT_IMPORT="github.com/briancain/gameboy-go/version"
LDFLAGS="-X $(GIT_IMPORT).Version=$(VERSION) -X $(GIT_IMPORT).Ref=$(REF)"
# END Builds

default: build

build:
	go build -ldflags $(LDFLAGS) -o bin/${BINARY} ./cmd/gameboy-go

run:
	./bin/${BINARY} $(ARGS)

clean:
	rm -rf bin/

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

all: clean build fmt vet test
