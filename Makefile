build:
	go build -v .

build-all: build-darwin build-linux

build-darwin:
	GOARCH=amd64 GOOS=darwin go build -v -o pkg/darwin/go-git

build-linux:
	GOARCH=amd64 GOOS=linux go build -v -o pkg/x86_64/go-git
