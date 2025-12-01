BINARY_NAME=confusio

all: clean build-all

build:
	go build -o $(BINARY_NAME) main.go

build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe main.go

clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*

.PHONY: all build build-all clean
