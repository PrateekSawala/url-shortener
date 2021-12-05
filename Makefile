# Build, test, package and more ... A developer everyday tool
BUILD_DIR := bin

default: all

all: clean build run

.PHONY: $(BUILD_DIR)/server
bin/server: cmd/*.go
	CGO_ENABLED=0 go build -o ./bin/server ./cmd/
	cp -R ./static ./bin/static
	cp -R ./templates ./bin/templates

.PHONY: build
build: clean bin/server

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: run
run: build
	bin/server

.PHONY: test
test:
	go test -v -cover -coverprofile=cover.out -covermode=atomic ./endpoint/... ./service/...

.PHONY: image
image: build
	docker build -t urlshortener:test -f Dockerfile ./bin/.