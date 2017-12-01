package user_interface

import (
    "fmt"
    "io"
    "os"
    "github.com/nettoclaudio/dontpad-cli/dontpad"
)

var (
    getContentFolder func(string) (dontpad.Response, error)
    getSubfolders    func(string) ([]string, error)
    outputChannel    io.Writer
)

func init() {
    getContentFolder = func(remoteFolder string) (dontpad.Response, error) { return  dontpad.GetContentFolder(remoteFolder) }

    getSubfolders = func(remoteFolder string) ([]string, error) { return dontpad.GetSubfolders(remoteFolder) }

    outputChannel = os.Stdout
}

func ShowContentFolder(remoteFolder string) error {
    response, err := getContentFolder(remoteFolder)

    if err == nil {
        fmt.Fprintf(outputChannel, "%s", response.Body)
    }

    return err
}

func ListSubfolders(remoteFolder string) error {
    subfolders, err := getSubfolders(remoteFolder)

    if err == nil {
        for _, subfolder := range subfolders {
            fmt.Fprintf(outputChannel, "%s\n", subfolder)
        }
    }

    return err
}
