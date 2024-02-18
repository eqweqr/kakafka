GO_VERSION = 1.20

.PHONY: install-go

setup: install-go 

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

build: 
	go build -o api cmd/main.go

	