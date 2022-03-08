NAMESPACE=briancain
NAME=gameboy-go
BINARY=${NAME}

default: build

build:
	go build -o bin/${BINARY}
