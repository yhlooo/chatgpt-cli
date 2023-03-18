GO_MODULE := github.com/keybrl/chatgpt-cli

.PHONY: build
build:
	go build -o bin/chatgpt-cli $(GO_MODULE)

.PHONY: build-all
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/chatgpt-cli $(GO_MODULE)
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/chatgpt-cli $(GO_MODULE)
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/chatgpt-cli $(GO_MODULE)
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/chatgpt-cli $(GO_MODULE)
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/chatgpt-cli $(GO_MODULE)
	GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/chatgpt-cli $(GO_MODULE)
