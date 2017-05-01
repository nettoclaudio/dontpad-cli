package main

import (
    "fmt"
    "os"
    "github.com/nettoclaudio/dontpad-cli/dontpad"
    "github.com/nettoclaudio/dontpad-cli/user_interface"
    "github.com/nettoclaudio/dontpad-cli/util"
)

func main() {
    var setup user_interface.SetUp
    var err   error

    setup, err = user_interface.ProcessCommands()

    util.ShowMessageAndExitOnError(err)

    var response dontpad.Response

    response, err = dontpad.GetContentFolder(setup.RemoteFolder)

    util.ShowMessageAndExitOnError(err)

    fmt.Fprintf(os.Stdout, "%s", response.Body)
}
