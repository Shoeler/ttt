GO ?= /usr/local/go/bin/go

.PHONY: build run test clean

build:
	$(GO) build -o ttt main.go

run:
	$(GO) run main.go

test:
	$(GO) test ./...

clean:
	rm -f ttt
