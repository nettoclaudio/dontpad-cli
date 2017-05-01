package util

import (
    "bytes"
    "errors"
    "testing"
)

func TestShowMessageAndExitOnError_ErrorNotFalse_MustShowErrorMessage(t *testing.T) {
    oldExit := exit
    oldErrorChannel := errorChannel

    defer func() {
        exit = oldExit
        errorChannel = oldErrorChannel
    }()

    exit = func(code int) {}

    errorBuffer := &bytes.Buffer{}
    errorChannel = errorBuffer

    expectedErrorMessage := "Something wrong occurred."
    err := errors.New(expectedErrorMessage)

    ShowMessageAndExitOnError(err)

    actualErrorMessage := errorBuffer.String()

    if actualErrorMessage != expectedErrorMessage {
        t.Errorf("Expected [%s] but got [%s].", expectedErrorMessage, actualErrorMessage)
    }
}

func TestShowMessageAndExitOnError_ErrorFalse_MustNotShowAnything(t *testing.T) {
    oldExit := exit
    oldErrorChannel := errorChannel

    defer func() {
        exit = oldExit
        errorChannel = oldErrorChannel
    }()

    exit = func(code int) {}

    errorBuffer := &bytes.Buffer{}
    errorChannel = errorBuffer

    ShowMessageAndExitOnError(nil)

    if errorBuffer.Len() != 0 {
        t.Errorf("Expected a false value but it didn't.")
    }
}

func TestShowMessageAndExitOnError_ErrorNotFalse_MustCallExitFunction(t *testing.T) {
    var wasExistCalled bool

    oldExit := exit
    oldErrorChannel := errorChannel

    defer func() {
        errorChannel = oldErrorChannel
        exit = oldExit
    }()

    exit = func(code int) { wasExistCalled = true }

    errorChannel = &bytes.Buffer{}

    ShowMessageAndExitOnError(errors.New("Error Message."))

    if ! wasExistCalled {
        t.Errorf("Expected true but got false.")
    }
}
