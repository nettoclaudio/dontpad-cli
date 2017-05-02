package util

import (
    "bytes"
    "errors"
    "io"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestExtractHttpResponseBodyIfStatusCodeIsOk_ConnectionRefusedError_MustReturnError(t *testing.T) {
    oldDoGetRequest := doGetRequest

    defer func() { doGetRequest = oldDoGetRequest }()

    doGetRequest = func(url string) (*http.Response, error) {
        return nil, errors.New("Connection refused.")
    }

    _, err := ExtractHttpResponseBodyIfStatusCodeIsOk("URL")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestExtractHttpResponseBodyIfStatusCodeIsOk_ResponseIsOk_MustReturnResponseBody(t *testing.T) {
    oldDoGetRequest := doGetRequest

    defer func() { doGetRequest = oldDoGetRequest }()

    expectedContent := "It works."

    doGetRequest = func(url string) (*http.Response, error) {
        response := &http.Response{
            Status:        "200 OK",
            StatusCode:    http.StatusOK,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          ioutil.NopCloser(bytes.NewBufferString(expectedContent)),
            ContentLength: int64(len(expectedContent)),
            Request:       nil,
            Header:        make(http.Header, 0),
        }

        return response, nil
    }

    actualContent, err := ExtractHttpResponseBodyIfStatusCodeIsOk("http://fakedlocalhost/")

    if err != nil {
        t.Errorf("Should not return an error.")
    }

    if string(actualContent) != expectedContent {
        t.Errorf("Expected [%s] but got [%s].", expectedContent, string(actualContent))
    }
}

func TestExtractHttpResponseBodyIfStatusCodeIsOk_PageNotFound_MustReturnAnError(t *testing.T) {
    oldDoGetRequest := doGetRequest

    defer func() { doGetRequest = oldDoGetRequest }()

    doGetRequest = func(url string) (*http.Response, error) {
        response := &http.Response{
            Status:        "404 NotFound",
            StatusCode:    http.StatusNotFound,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          ioutil.NopCloser(bytes.NewBufferString("")),
            ContentLength: int64(len("")),
            Request:       nil,
            Header:        make(http.Header, 0),
        }

        return response, nil
    }

    _, err := ExtractHttpResponseBodyIfStatusCodeIsOk("http://fakedlocalhost/")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}

func TestExtractHttpResponseBodyIfStatusCodeIsOk_ResponseIsOk_ProblemOnReader_MustReturnAnError(t *testing.T) {
    oldDoGetRequest := doGetRequest
    oldReadAll := readAll

    defer func() {
        doGetRequest = oldDoGetRequest
        readAll = oldReadAll
    }()

    content := "It works."

    doGetRequest = func(url string) (*http.Response, error) {
        response := &http.Response{
            Status:        "200 OK",
            StatusCode:    http.StatusOK,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          ioutil.NopCloser(bytes.NewBufferString(content)),
            ContentLength: int64(len(content)),
            Request:       nil,
            Header:        make(http.Header, 0),
        }

        return response, nil
    }

    readAll = func(r io.Reader) ([]byte, error) { return []byte{}, errors.New("Problem on EOF") }

    _, err := ExtractHttpResponseBodyIfStatusCodeIsOk("http://fakedlocalhost/")

    if err == nil {
        t.Errorf("Should return an error.")
    }
}
