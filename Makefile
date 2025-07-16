NAMESPACE=briancain
NAME=gameboy-go
BINARY=${NAME}

default: build

build:
	go build -o bin/${BINARY}

run:
	./bin/${BINARY} $(ARGS)

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

all: clean build fmt vet test
