GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
DISTDIR=dist
BINARY_NAME=content-management-api
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m
COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

all: clean install fmt test build

build:
	$(GOBUILD) -o $(DISTDIR)/$(BINARY_NAME) -v

fmt:
	$(GOFMT) ./...

test:
	$(GOTEST) -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GORUN) main.go

binary-run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

install:
	go mod download && go mod tidy

cross-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -tags netgo -installsuffix netgo -o $(DISTDIR)/$(BINARY_NAME)-linux-amd64 -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -a -tags netgo -installsuffix netgo -o $(DISTDIR)/$(BINARY_NAME)-darwin-amd64 -v
