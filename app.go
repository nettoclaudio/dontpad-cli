package main

import (
    "github.com/nettoclaudio/dontpad-cli/user_interface"
    "github.com/nettoclaudio/dontpad-cli/util"
)

func main() {
    var setup user_interface.SetUp
    var err   error

    setup, err = user_interface.ProcessCommands()

    util.ShowMessageAndExitOnError(err)

    if setup.ListSubfolders {
        err = user_interface.ListSubfolders(setup.RemoteFolder)

        util.ShowMessageAndExitOnError(err)

        return
    }

    if user_interface.HasPipedInput() {
        inputData := user_interface.GetPipedInputData()

        user_interface.WriteFolder(setup.RemoteFolder, inputData);
    } else {
        err = user_interface.ShowContentFolder(setup.RemoteFolder)
        
        util.ShowMessageAndExitOnError(err)
    }
}
