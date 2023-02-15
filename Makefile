default: build

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

build:
	go build -o build/$(GOOS)/$(GOARCH)/lift

build-all: build/linux/amd64/lift build/darwin/amd64/lift

build/%/amd64/lift:
	GOOS=$* GOARCH=amd64 go build -v -o $@ .