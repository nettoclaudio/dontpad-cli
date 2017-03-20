package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "github.com/nettoclaudio/dontpad-cli/html/util"
)

func main() {
    
    if ! hasFolderArgument() {
        showProgramUsage()
        
        return
    }
    
    folder := os.Args[1]
    
    if ! isValidFolder(folder) {
        fmt.Fprintf(os.Stderr, "ERROR: That is not a valid folder.\n")
        
        return
    }
    
    finalURL := formatDontpadURLWithFolder(folder)
    
    content, err := getFolderMessage(finalURL)
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: Problem getting message.\n%s", err)
    }
    
    fmt.Printf("%s", content)
}

func hasFolderArgument() bool {
    
    if len(os.Args) < 2 {
        return false
    }
    
    return true
}

func showProgramUsage() {
    fmt.Fprintf(os.Stderr, "Usage: %s FOLDER\n", os.Args[0])
}

func isValidFolder(folder string) bool {
    
    if folder == "" || folder == "/" {
        return false
    }
    
    return true
}

func formatDontpadURLWithFolder(folder string) string {
    return "http://dontpad.com/" + folder
}

func getFolderMessage(url string) (string, error) {
    
    response, err := http.Get(url)
    
    if err != nil {
        return "", err
    }
    
    defer response.Body.Close()
    
    page, err := ioutil.ReadAll(response.Body)
    
    if err != nil {
        return "", err
    }
    
    content, err := util.ExtractFolderMessage(string(page))
    
    if err != nil {
        return "", err
    }

    return content, nil
}
