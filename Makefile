
.DEFAULT_GOAL := build

fmt:
	gofmt -w -s .
.PHONY:fmt

lint: fmt
	golint .
.PHONY:lint

vet: fmt
	go vet ./...
	# shadow ./...
.PHONY:vet

build: vet
	go build -o bin/c9s cmd/c9s/main.go
.PHONY:build

test:
	go test
.PHONY:test

clean:
	$(RM) -rf bin/
.PHONY:clean
