package util

import (
    "fmt"
    "io"
    "os"
)

const (
    StatusCodeSuccess = iota
    StatusCodeError
)

var (
    exit           func(int)
    errorChannel   io.Writer
)

func init() {
    exit = func(code int) { os.Exit(code) }

    errorChannel = os.Stderr
}

func ShowMessageAndExitOnError(err error) {
    if err != nil {
        fmt.Fprintf(errorChannel, "%s", err)

        exit(StatusCodeError)
    }
}
