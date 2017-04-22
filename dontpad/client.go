package dontpad

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Response struct {
    Changed        bool     `json:"changed"`
    LastUpdate     int      `json:"lastUpdate"`
    Body           string   `json:"body"`
}

var (
    doGetRequest func(string) (*http.Response, error)
    templateURLViewFolder string = "http://dontpad.com/%s.body.json?lastUpdate=%d"
)

func init() {
    doGetRequest = func(url string) (*http.Response, error) {return http.Get(url)}
}

func GetContentFolder(remoteFolder string) (Response, error) {
    var response Response

    url := formatTemplateURLViewFolder(remoteFolder, 0)

    resp, err := doGetRequest(url)

    if err != nil {
        return response, err
    }

    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        content, _ := ioutil.ReadAll(resp.Body)

        json.Unmarshal(content, &response)

        return response, nil
    }

    return response, errors.New("Unexpected response(status code).")
}

func formatTemplateURLViewFolder(remoteFolder string, lastUpdate int) string {
    buffer := &bytes.Buffer{}

    fmt.Fprintf(buffer, templateURLViewFolder, remoteFolder, 0)

    return buffer.String()
}
