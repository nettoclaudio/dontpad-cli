# dontpad-cli

A minimal tool for Dontpad's users under CLI.

**Warning**: This app isn't a Dontpad official software.

[![CircleCI Status](https://circleci.com/gh/nettoclaudio/dontpad-cli.svg?style=svg&circle-token=3bd4b3d74f5089c30aa224545365a5585e6c994d)](https://circleci.com/gh/nettoclaudio/dontpad-cli)
[![Coverage Status](https://coveralls.io/repos/github/nettoclaudio/dontpad-cli/badge.svg?branch=master)](https://coveralls.io/github/nettoclaudio/dontpad-cli?branch=master)
[![MicroBadger Info](https://images.microbadger.com/badges/image/nettoclaudio/dontpad-cli.svg)](https://microbadger.com/images/nettoclaudio/dontpad-cli)

--- 

## Usage

[![asciicast](https://asciinema.org/a/150757.png)](https://asciinema.org/a/150757)

## Quick start

```bash
go get github.com/nettoclaudio/dontpad-cli
${GOPATH:-"~/go"}/bin/dontpad-cli /my-first-folder/annotations
```

or (preferred for developers)

```bash
mkdir -p ${GOPATH}/src/github.com/nettoclaudio
cd ${GOPATH}/src/github.com/nettoclaudio
git clone https://github.com/nettoclaudio/dontpad-cli.git
cd dontpad-cli
make
./bin/dontpad-cli /my-first-folder/annotations

# Optional - Put the executable in your PATH
sudo make install
dontpad-cli /my-first-folder/annotations
```

## TO DO

+   ~~View a folder~~
+   ~~Edit a folder~~
+   ~~List subfolders~~
+   Backup a folder(and subfolders)

Made with ~~Dontpad~~ <3
