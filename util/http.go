package util

import (
    "errors"
    "net/http"
    "io"
    "io/ioutil"
)

var (
    doGetRequest   func(string) (*http.Response, error)
    readAll        func(io.Reader) ([]byte, error)
)

func init() {
    doGetRequest = func(url string) (*http.Response, error) { return http.Get(url) }
    readAll      = func(r io.Reader) ([]byte, error) { return ioutil.ReadAll(r) }
}

func ExtractHttpResponseBodyIfStatusCodeIsOk(url string) ([]byte, error) {
    var body []byte

    response, err := doGetRequest(url)

    if err != nil {
        return body, err
    }

    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        body, err = readAll(response.Body)

        return body, err
    } else {
        return body, errors.New("Unexpected status code.")
    }
}
