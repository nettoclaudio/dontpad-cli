package main

import (
    "fmt"
    "os"
    "github.com/nettoclaudio/dontpad-cli/user_interface"
    "github.com/nettoclaudio/dontpad-cli/dontpad"
)

func main() {
    var setup user_interface.SetUp
    var err   error

    setup, err = user_interface.ProcessCommands()

    showAndExitOnError(err)

    var response dontpad.Response

    response, err = dontpad.GetContentFolder(setup.RemoteFolder)

    showAndExitOnError(err)

    fmt.Fprintf(os.Stdout, "%s", response.Body)
}

func showAndExitOnError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s", err)

        os.Exit(1)
    }
}
