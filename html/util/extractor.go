package util

import (
    "strings"
    "golang.org/x/net/html"
)

func ExtractFolderMessage(htmlPage string) (string, error) {
    
    documentNode, err := html.Parse(strings.NewReader(htmlPage))
    
    if err != nil {
        return "", err
    }
    
    documentCNode := CNode{documentNode}
    
    folderNode, err := documentCNode.GetElementByID("text")
    
    if err != nil {
        return "", err
    }
    
    if folderNode.FirstChild == nil {
        return "", nil
    }
    
    message := folderNode.FirstChild.Data
    
    return message, nil
}
