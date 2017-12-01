package user_interface

import (
    "bytes"
    "errors"
    "testing"
    "github.com/nettoclaudio/dontpad-cli/dontpad"
)

func TestShowContentFolder_NoError_MustPrintResponseBodyOnBuffer(t *testing.T) {
    oldGetContentFolder := getContentFolder
    oldOutputChannel := outputChannel

    defer func() {
        getContentFolder = oldGetContentFolder
        outputChannel = oldOutputChannel
    }()

    expected := "Hey Dude.\n How are you?"

    getContentFolder = func(remoteFolder string) (dontpad.Response, error) { return  dontpad.Response{Body: expected}, nil }

    buffer := &bytes.Buffer{}
    outputChannel = buffer

    ShowContentFolder("my-folder")

    actual := buffer.String()

    if actual != expected {
        t.Errorf("Expected [%s] but got [%s].",  expected, actual)
    }
}

func TestShowContentFolder_TimeoutError_MustNotPrintAnything(t *testing.T) {
    oldGetContentFolder := getContentFolder
    oldOutputChannel := outputChannel

    defer func() {
        getContentFolder = oldGetContentFolder
        outputChannel = oldOutputChannel
    }()

    getContentFolder = func(remoteFolder string) (dontpad.Response, error) { return  dontpad.Response{}, errors.New("Timeout") }
    buffer := &bytes.Buffer{}

    outputChannel = buffer

    ShowContentFolder("my-folder")

    if buffer.Len() > 0 {
        t.Errorf("Should not print on stadard output.")
    }
}

func TestListSubfolders_NoError_MustPrintOneSubfolderPerLine(t *testing.T) {
    oldGetSubfolders := getSubfolders
    oldOutputChannel := outputChannel

    defer func() {
        getSubfolders = oldGetSubfolders
        outputChannel = oldOutputChannel
    }()

    getSubfolders = func(remoteFolder string) ([]string, error) { return []string{"sub1", "sub2", "sub3"}, nil }

    buffer := &bytes.Buffer{}
    outputChannel = buffer

    ListSubfolders("my-folder")

    actual := buffer.String()

    expected := "sub1\nsub2\nsub3\n"

    if actual != expected {
        t.Errorf("Expected [%s] but got [%s].",  expected, actual)
    }
}

func TestListSubfolders_TimeoutErrors_MustNotPrintAnything(t *testing.T) {
    oldGetSubfolders := getSubfolders
    oldOutputChannel := outputChannel

    defer func() {
        getSubfolders = oldGetSubfolders
        outputChannel = oldOutputChannel
    }()

    getSubfolders = func(remoteFolder string) ([]string, error) { return []string{}, errors.New("Timeout") }

    buffer := &bytes.Buffer{}
    outputChannel = buffer

    ListSubfolders("my-folder")

    if buffer.Len() > 0 {
        t.Errorf("Should not print on stadard output.")
    }
}
