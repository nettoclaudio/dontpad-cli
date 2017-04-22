package dontpad

import (
    "bytes"
    "errors"
    "io/ioutil"
    "net/http"
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
    oldDoGetRequest := doGetRequest

    doGetRequest = func(url string) (*http.Response, error) {
        return nil, errors.New("Timeout")
    }

    defer func() { doGetRequest = oldDoGetRequest }()

    _, err := GetContentFolder("my-remote-folder")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestGetContentFolder_UnexpectedResponse_MustReturnError(t *testing.T) {
    oldDoGetRequest := doGetRequest

    doGetRequest = func(url string) (*http.Response, error) {

        response := &http.Response{
            Status:        "404 Not Found",
            StatusCode:    http.StatusNotFound,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          ioutil.NopCloser(bytes.NewBufferString("")),
            ContentLength: int64(0),
            Request:       nil,
            Header:        make(http.Header, 0),
        }

        return response, nil
    }

    defer func() { doGetRequest = oldDoGetRequest }()

    _, err := GetContentFolder("my-remote-folder")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestGetContentFolder_ResponseOK_MustNotReturnError(t *testing.T) {
    oldDoGetRequest := doGetRequest

    doGetRequest = func(url string) (*http.Response, error) {

        json := `{"changed":false,"lastUpdate":0,"body":"Hello, howdy!"}`

        response := &http.Response{
            Status:        "200 OK",
            StatusCode:    http.StatusOK,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          ioutil.NopCloser(bytes.NewBufferString(json)),
            ContentLength: int64(len(json)),
            Request:       nil,
            Header:        make(http.Header, 0),
        }

        return response, nil
    }

    defer func() { doGetRequest = oldDoGetRequest }()

    resp, err := GetContentFolder("my-remote-folder")

    if err != nil {
        t.Errorf("Should not return an error.")
    }

    expectedContent := "Hello, howdy!"

    if resp.Body != expectedContent {
        t.Errorf("Expected [%s] but got [%s].", expectedContent, resp.Body)
    }
}
