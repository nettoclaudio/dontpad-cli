FROM golang:alpine

LABEL maintainer="Claudio Netto <nettoclaudio@ufrj.br>"

RUN apk update && apk add git && \
    go get github.com/nettoclaudio/dontpad-cli
