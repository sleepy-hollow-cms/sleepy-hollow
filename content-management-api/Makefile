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
IS_DOCKER_INSTALLED=$(shell which docker >> /dev/null 2>&1; echo $$?)
DOCKER_DIR=docker

.PHONY: all
all: clean install fmt test build

.PHONY: build
build:
	$(GOBUILD) -o $(DISTDIR)/$(BINARY_NAME) -v

.PHONY: fmt
fmt:
	$(GOFMT) ./...

.PHONY: test
test:
	$(GOTEST) -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(DISTDIR)/$(BINARY_NAME)*

.PHONY: run
run: export APP_PORT=3000
run: export MONGO_USER=root
run: export MONGO_PASSWORD=password
run: export MONGO_HOST=localhost
run: export MONGO_PORT=27017
run:
	$(GORUN) main.go

.PHONY: install
install:
	go mod download && go mod tidy

.PHONY: build
cross-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -tags netgo -installsuffix netgo -o $(DISTDIR)/$(BINARY_NAME)-linux-amd64 -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -a -tags netgo -installsuffix netgo -o $(DISTDIR)/$(BINARY_NAME)-darwin-amd64 -v

.PHONY: check-docker
check-docker:
	@if [ $(IS_DOCKER_INSTALLED) -eq 1 ]; \
		then \
		echo "ERROR:\tdocker is not installed!" \
		&& exit 1; \
	fi;

.PHONY: docker-build
docker-build: check-docker
	@docker build -t content-management-api -f $(DOCKER_DIR)/Dockerfile .

.PHONY: compose-up
compose-up: docker-build
	@docker-compose -f $(DOCKER_DIR)/docker-compose.yml up -d

.PHONY: compose-down
compose-down:
	@docker-compose -f $(DOCKER_DIR)/docker-compose.yml down
