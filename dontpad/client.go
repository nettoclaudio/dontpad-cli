package dontpad

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "github.com/nettoclaudio/dontpad-cli/util"
)

type Response struct {
    Changed        bool     `json:"changed"`
    LastUpdate     int      `json:"lastUpdate"`
    Body           string   `json:"body"`
}

var (
    extractHttpResponseBody func(string) ([]byte, error)

    templateURL           string = "http://dontpad.com/%s"
    templateURLViewFolder string =  templateURL + ".body.json?lastUpdate=%d"
    templateURLSubfolders string =  templateURL + ".menu.json"
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

func GetSubfolders(remoteFolder string) ([]string, error) {
    var subfolders []string

    url := formatTemplateURLSubfolders(remoteFolder)

    body, err := extractHttpResponseBody(url)

    if err == nil {
        json.Unmarshal(body, &subfolders)
    }

    return subfolders, err
}

func ReplaceContentFolder(remoteFolder, data string) {
    finalURL := formatTemplateURL(remoteFolder)

    formValues := url.Values{}
    formValues.Set("text", data)

    http.PostForm(finalURL, formValues)
}

func formatTemplateURL(remoteFolder string) string {
    buffer := &bytes.Buffer{}
    
    fmt.Fprintf(buffer, templateURL, remoteFolder)

    return buffer.String()
}

func formatTemplateURLViewFolder(remoteFolder string, lastUpdate int) string {
    buffer := &bytes.Buffer{}

    fmt.Fprintf(buffer, templateURLViewFolder, remoteFolder, 0)

    return buffer.String()
}

func formatTemplateURLSubfolders(remoteFolder string) string {
    buffer := &bytes.Buffer{}

    fmt.Fprintf(buffer, templateURLSubfolders, remoteFolder)

    return buffer.String()
}
