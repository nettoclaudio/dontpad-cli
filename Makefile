GO ?= go

goBuildFlags ?= -race -x
goBuildDir ?= bin
goAppName ?= dontpad-cli

installDir ?= /usr/local/bin

.PHONY: all build pre-build clean install uninstall test

all: build

build: pre-build
	$(GO) build $(goBuildFlags) -o $(goBuildDir)/$(goAppName) app.go

pre-build: clean
	mkdir $(goBuildDir)

clean:
	rm -f $(goBuildDir)/$(goAppName)
	rm -df $(goBuildDir)

install:
	ln -s $(shell pwd)/$(goBuildDir)/$(goAppName) $(installDir)

uninstall:
	rm -f $(installDir)/$(goAppName)

test:
	$(GO) test -v ./...
