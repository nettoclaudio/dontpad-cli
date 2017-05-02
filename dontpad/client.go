package dontpad

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/nettoclaudio/dontpad-cli/util"
)

type Response struct {
    Changed        bool     `json:"changed"`
    LastUpdate     int      `json:"lastUpdate"`
    Body           string   `json:"body"`
}

var (
    extractHttpResponseBody func(string) ([]byte, error)
    templateURLViewFolder string = "http://dontpad.com/%s.body.json?lastUpdate=%d"
)

func init() {
    extractHttpResponseBody = func(url string) ([]byte, error) { return util.ExtractHttpResponseBodyIfStatusCodeIsOk(url) }
}

func GetContentFolder(remoteFolder string) (Response, error) {
    var response Response

    url := formatTemplateURLViewFolder(remoteFolder, 0)

    body, err := extractHttpResponseBody(url)

    if err == nil {
        json.Unmarshal(body, &response)
    }

    return response, err
}

func formatTemplateURLViewFolder(remoteFolder string, lastUpdate int) string {
    buffer := &bytes.Buffer{}

    fmt.Fprintf(buffer, templateURLViewFolder, remoteFolder, 0)

    return buffer.String()
}
