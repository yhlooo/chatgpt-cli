GO_MODULE := github.com/keybrl/chatgpt-cli
VERSION ?= 0.0.0-dev

.PHONY: build
build:
	go build -o bin/chatgpt-cli $(GO_MODULE)

.PHONY: build-all
build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/chatgpt-cli $(GO_MODULE)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/chatgpt-cli $(GO_MODULE)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/chatgpt-cli $(GO_MODULE)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/chatgpt-cli $(GO_MODULE)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/chatgpt-cli $(GO_MODULE)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/chatgpt-cli $(GO_MODULE)

.PHONY: release
release: build-all
	mkdir -p bin/archive
	cp bin/linux/amd64/chatgpt-cli bin/archive/chatgpt-cli-linux-amd64
	cp bin/linux/arm64/chatgpt-cli bin/archive/chatgpt-cli-linux-arm64
	cp bin/darwin/amd64/chatgpt-cli bin/archive/chatgpt-cli-darwin-amd64
	cp bin/darwin/arm64/chatgpt-cli bin/archive/chatgpt-cli-darwin-arm64
	cp bin/windows/amd64/chatgpt-cli bin/archive/chatgpt-cli-windows-amd64
	cp bin/windows/arm64/chatgpt-cli bin/archive/chatgpt-cli-windows-arm64
