package user_interface

import (
    "bytes"
    "flag"
    "os"
    "testing"
)

func TestCustomUsage_WriteUsageMessageInBuffer_MustWriteSameUsageMessage(t *testing.T) {
    buffer := &bytes.Buffer{}

    outputDefault = buffer
    programName = "dontpad-cli.exe"

    customUsage()

    actualHeader := buffer.String()
    expectedHeader := "Usage:  dontpad-cli.exe <remote-folder>\n"

    if actualHeader != expectedHeader {
        t.Errorf("Expected [%s] but got [%s].", expectedHeader, actualHeader)
    }
}

func TestHasRemoteFolder_NoArgs_MustReturnFalse(t *testing.T) {
    os.Args = []string{"dontpad-cli"}

    flag.Parse()

    if hasRemoteFolder() == true {
        t.Errorf("Expected [false] but got [true].")
    }
}

func TestHasRemoteFolder_ContainsRemoteFolderArg_MustReturnTrue(t *testing.T) {
    os.Args = []string{"dontpad-cli", "/my-remote-folder/subfolder"}

    flag.Parse()

    if hasRemoteFolder() == false {
        t.Errorf("Expected [true] but got [false].")
    }
}

func TestHasRemoteFolder_ContainsHelpFlag_MustReturnFalse(t *testing.T) {
    os.Args = []string{"dontpad-cli", "--help"}

    flag.CommandLine = flag.NewFlagSet(os.Args[0],  flag.ContinueOnError)

    flag.Parse()

    if hasRemoteFolder() == true {
        t.Errorf("Expected [false] but got [true].")
    }
}

func TestHasRemoteFolder_ContainsHelpFlag_ContainsRemoteFolder_MustReturnTrue(t *testing.T) {
    os.Args = []string{"dontpad-cli", "--help", "/my-remote-folder"}

    flag.CommandLine = flag.NewFlagSet(os.Args[0],  flag.ContinueOnError)

    flag.Parse()

    if hasRemoteFolder() == false {
        t.Errorf("Expected [true] but got [false].")
    }
}

func TestSanitizeRemoteFolder_RemoteFolderUnsanitized_MustReturnDifferentRemoteFolderFromOriginal(t *testing.T) {
    expectedRemoteFolder := "my-folder"

    actualSanitizedRemoteFolder :=  sanitizeRemoteFolder(" /my-folder ")

    if actualSanitizedRemoteFolder != expectedRemoteFolder {
        t.Errorf("Expected [%s] but got [%s].", expectedRemoteFolder, actualSanitizedRemoteFolder)
    }
}

func TestSanitizeRemoteFolder_RemoteFolderSanitized_MustReturnSameRemoteFolderFromOriginal(t *testing.T) {
    expectedRemoteFolder := "my-folder/subfolder"

    actualSanitizedRemoteFolder :=  sanitizeRemoteFolder("my-folder/subfolder")

    if actualSanitizedRemoteFolder != expectedRemoteFolder {
        t.Errorf("Expected [%s] but got [%s].", expectedRemoteFolder, actualSanitizedRemoteFolder)
    }
}

func TestIsValidRemoteFolder_InvalidsPrefixes_AllMustReturnFalse(t *testing.T) {
    invalidFolders := []string{"", "/", "/folder.zip", "/static/forbidden/folder"}

    for _, folder := range invalidFolders {
        if isValidRemoteFolder(folder) {
            t.Errorf("Folder [%s] is invalid but got valid.", folder)
        }
    }
}

func TestIsValidRemoteFolder_ValidPrefixes_AllMustReturnTrue(t *testing.T) {
    validFolders := []string{"my-folder", "/my/subfolder", "test.zipa"}

    for _, folder := range validFolders {
        if ! isValidRemoteFolder(folder) {
            t.Errorf("Folder [%s] is valid but got invalid.", folder)
        }
    }
}

func TestProcessCommands_NoRemoteFolder_MustReturnError(t *testing.T) {
    os.Args = []string{"dontpad-cli"}

    flag.Parse()

    _, err := ProcessCommands()

    if err == nil {
        t.Errorf("Expected not nil error.")
    }
}

func TestProcessCommands_InvalidRemotePath_MustReturnError(t *testing.T) {
    os.Args = []string{"dontpad-cli", "/static/invalid/path"}

    flag.Parse()

    _, err := ProcessCommands()

    if err == nil {
        t.Errorf("Expected not nil error.")
    }
}

func TestProcessCommands_MustReturnSetupCorrectly(t *testing.T) {
    os.Args = []string{"dontpad-cli", "/my-folder/agenda"}

    flag.Parse()

    setup, err := ProcessCommands()

    if err != nil {
        t.Errorf("Expected nil error.")
    }

    if setup.RemoteFolder != "my-folder/agenda" {
        t.Errorf("Expected [my-folder/agenda] but got [%s]", setup.RemoteFolder)
    }
}

func TestHasPipedInput_ReceivedPipedInput_MustReturnTrue(t *testing.T) {
    oldGetInputFileMode := getInputFileMode
    
    defer func() { getInputFileMode = oldGetInputFileMode }()

    getInputFileMode = func() os.FileMode {
        return os.ModeNamedPipe
    }

    if ! HasPipedInput() {
        t.Errorf("Expected [true] but got [false]")
    }
}

func TestHasPipedInput_NotReceivedPipedInput_MustReturnFalse(t *testing.T) {
    oldGetInputFileMode := getInputFileMode
    
    defer func() { getInputFileMode = oldGetInputFileMode }()

    getInputFileMode = func() os.FileMode {
        return os.ModeDir
    }

    if HasPipedInput() {
        t.Errorf("Expected [false] but got [true]")
    }
}
