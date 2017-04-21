GO ?= go

goBuildFlags ?= -race -x
goBuildDir ?= bin
goAppName ?= dontpad-cli

installDir ?= /usr/local/bin
reportDir ?= report

packages := $(shell $(GO) list ./...)

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

test: pre-test
	echo "mode: count" > $(reportDir)/coverage-all.out
	$(foreach pkg,$(packages),\
		touch $(reportDir)/coverage.out;\
		$(GO) test -coverprofile=$(reportDir)/coverage.out -covermode=count $(pkg);\
		tail -n +2 $(reportDir)/coverage.out >> $(reportDir)/coverage-all.out;\
	)
	mv -f $(reportDir)/coverage-all.out $(reportDir)/coverage.out

pre-test: clean-old-reports
	mkdir -p $(reportDir)

clean-old-reports:
	rm -rf $(reportDir)
