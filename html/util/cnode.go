package util

import (
    "errors"
    "strings"
    "golang.org/x/net/html"
)

// Custom Node implements special methods
type CNode struct {
    *html.Node
}

func (node *CNode) GetElementByID(id string) (*CNode, error) {
    
    nodeFound := node.reachElementID(id)
    
    if nodeFound == nil {
        return nil, errors.New("Element ID not found.")
    }
    
    return nodeFound, nil
}

func (node *CNode) reachElementID(id string) *CNode {
    
    if node.hasElementID(id) {
        return node
    }
    
    var nextNode *CNode

    if node.FirstChild != nil {
        nextNode = &CNode{node.FirstChild}
    }
    
    for nextNode != nil {
        
        foundNode := nextNode.reachElementID(id)
        
        if foundNode != nil {
            return foundNode
        }
        
        if nextNode.NextSibling != nil {
            nextNode = &CNode{nextNode.NextSibling}
        } else {
            nextNode = nil
        }
    }
    
    return nil
}

func (node *CNode) hasElementID(id string) bool {
    
    if node.Type == html.ElementNode {
        for _, attribute := range node.Attr {
            if strings.EqualFold(attribute.Key, "id") {
                return strings.EqualFold(attribute.Val, id)
            }
        }
    }
    
    return false
}
