NAMESPACE=briancain
NAME=gameboy-go
BINARY=${NAME}
VERSION_FILE=version.txt
VERSION:=$(shell cat $(VERSION_FILE))

default: build

build:
	go build -o bin/${BINARY}
