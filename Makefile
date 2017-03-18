GO ?= go

goBuildFlags ?= -race -x
goBuildDir ?= bin
goAppName ?= dontpad-cli

installDir ?= /usr/local/bin

.PHONY: all build pre-build clean

all: build

build: pre-build
	$(GO) build $(goBuildFlags) -o $(goBuildDir)/$(goAppName) app.go

pre-build: clean
	mkdir $(goBuildDir)

clean: uninstall
	rm -f $(goBuildDir)/$(goAppName)
	rm -df $(goBuildDir)

install:
	ln -s $(shell pwd)/$(goBuildDir)/$(goAppName) $(installDir)

uninstall:
	rm -f $(installDir)/$(goAppName)
