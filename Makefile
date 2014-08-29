build:
	go build -v .

build-darwin:
	GOARCH=amd64 GOOS=darwin go build -v -o go-git_darwin

build-linux:
	GOARCH=amd64 GOOS=linux go build -v -o go-git_x86_64
