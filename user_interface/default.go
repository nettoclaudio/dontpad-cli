package user_interface

import (
    "errors"
    "flag"
    "fmt"
    "io"
    "os"
    "strings"
)

type SetUp struct {
    RemoteFolder   string
    ListSubfolders bool
}

var (
    programName    string
    outputDefault  io.Writer
    setup          SetUp
)

func init() {
    programName = os.Args[0]
    outputDefault = os.Stderr

    flag.Usage = customUsage

    flag.BoolVar(&setup.ListSubfolders, "subfolders", false, "List all subfolders from a remote folder.")

    flag.Parse()
}

func ProcessCommands() (SetUp, error) {
    if ! hasRemoteFolder() {
        flag.Usage()

        return setup, errors.New("Remote folder is required.")
    }

    setup.RemoteFolder = sanitizeRemoteFolder(flag.Arg(0))

    if ! isValidRemoteFolder(setup.RemoteFolder) {
        return setup, errors.New("[ERROR] The remote folder '" + setup.RemoteFolder +"' is invalid.")
    }

    return setup, nil
}

func customUsage() {
    usageHeader := "Usage:  %s <remote-folder>\n"

    fmt.Fprintf(outputDefault, usageHeader, programName)

    flag.PrintDefaults()
}

func hasRemoteFolder() bool {

    numberOfRemainingArgs := flag.NArg()

    if numberOfRemainingArgs > 0 {
        return true
    }

    return false
}

func sanitizeRemoteFolder(remoteFolder string) string {
    return strings.Trim(remoteFolder, " /")
}

func isValidRemoteFolder(remoteFolder string) bool {

    sanitizedRemoteFolder := sanitizeRemoteFolder(remoteFolder)

    if strings.HasPrefix(sanitizedRemoteFolder, "static/") ||
       strings.HasSuffix(sanitizedRemoteFolder, ".zip") ||
       sanitizedRemoteFolder == "" {
        return false
    }

    return true
}
