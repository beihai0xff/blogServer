# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=build/blogServer
BINARY_NAME_TEST=build/test


all: clean build_linux
deploy: all deployment
build_linux:
		CGO_ENABLED=0
		GOOS=linux
		GOARCH=amd64
		$(GOBUILD) -v -o $(BINARY_NAME) -tags=jsoniter .

build_darwin:
		CGO_ENABLED=0
		GOOS=darwin
		GOARCH=amd64
		$(GOBUILD) -v -o $(BINARY_NAME) -tags=jsoniter .

build_test:
		CGO_ENABLED=0
		GOOS=linux
		GOARCH=amd64
		$(GOBUILD) -v -o $(BINARY_NAME_TEST) -tags=jsoniter .

deployment:
		cp $(BINARY_NAME) ~/
		cd ~/
		nohup ./blogServer

test:
		$(GOTEST) -v ./...
bench:
		$(GOTEST) -bench=. -benchtime=3s -run=none
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)

run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

push:
		git push -u origin master
