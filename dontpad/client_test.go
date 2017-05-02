package dontpad

import (
    "errors"
    "testing"
)

func TestFormatTemplateURLViewFolder_MustReturnSameStringAsExpected(t *testing.T) {
    remoteFolder := "my-folder"
    lastUpdate := 0

    expectedURL := "http://dontpad.com/my-folder.body.json?lastUpdate=0"

    actualURL := formatTemplateURLViewFolder(remoteFolder, lastUpdate)

    if actualURL != expectedURL {
        t.Errorf("Expected [%s] but got [%s].", expectedURL, actualURL)
    }
}

func TestGetContentFolder_ConnectionTimeout_MustReturnError(t *testing.T) {
    oldExtractHttpResponseBody := extractHttpResponseBody

    defer func() { extractHttpResponseBody = oldExtractHttpResponseBody }()

    extractHttpResponseBody = func(url string) ([]byte, error) {
        return []byte{}, errors.New("Timeout")
    }

    _, err := GetContentFolder("my-remote-folder")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestGetContentFolder_UnexpectedResponse_MustReturnError(t *testing.T) {
    oldExtractHttpResponseBody := extractHttpResponseBody

    defer func() { extractHttpResponseBody = oldExtractHttpResponseBody }()

    extractHttpResponseBody = func(url string) ([]byte, error) {
        return []byte{}, errors.New("Unexpected status code.")
    }

    _, err := GetContentFolder("my-remote-folder")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestGetContentFolder_ResponseOK_MustNotReturnError_MustReturnSameBodyMessage(t *testing.T) {
    oldExtractHttpResponseBody := extractHttpResponseBody

    defer func() { extractHttpResponseBody = oldExtractHttpResponseBody }()

    extractHttpResponseBody = func(url string) ([]byte, error) {
        return []byte(`{"changed":false,"lastUpdate":0,"body":"Hello, howdy!"}`), nil
    }

    resp, err := GetContentFolder("my-remote-folder")

    if err != nil {
        t.Errorf("Should not return an error.")
    }

    expectedContent := "Hello, howdy!"

    if resp.Body != expectedContent {
        t.Errorf("Expected [%s] but got [%s].", expectedContent, resp.Body)
    }
}
