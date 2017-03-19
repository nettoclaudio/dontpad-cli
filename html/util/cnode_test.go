package util

import (
    "strings"
    "testing"
    "golang.org/x/net/html"
)

func TestHasElementID_ElementWithouIDAttribute_MustReturnFalse(t *testing.T) {
    element := `<div>I have not attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    divCNode := CNode{divElementNode}
    
    if divCNode.hasElementID("any-div-id") {
        t.Errorf("Should return false.")
    }
}

func TestHasElementID_ElementWithIDAttribute_MustReturnTrue(t *testing.T) {
    element := `<div id="content">I have attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    divCNode := CNode{divElementNode}
    
    if ! divCNode.hasElementID("content") {
        t.Errorf("Should return true.")
    }
}

func TestHasElementID_ElementWithIDAttributeButIDIsNotSame_MustReturnFalse(t *testing.T) {
    element := `<div id="another-content">I have attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    divCNode := CNode{divElementNode}
    
    if divCNode.hasElementID("content") {
        t.Errorf("Should return false.")
    }
}

func TestHasElementID_ElementWithIDAttributeAndIDUppercase_MustReturnTrue(t *testing.T) {
    element := `<div ID="CONTENT">I have attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    divCNode := CNode{divElementNode}
    
    if ! divCNode.hasElementID("content") {
        t.Errorf("Should return true.")
    }
}

func TestReachElementID_ElementIDExists_MustReturnSamePointerAddress(t *testing.T) {
    element := `<div id="content">I have attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    docCNode := CNode{doc}
    
    addressNodeExpected := divElementNode
    addressNodeGot := docCNode.reachElementID("content").Node

    if addressNodeExpected != addressNodeGot {
        t.Errorf("Pointer address expected [%p] but got [%p].", addressNodeExpected, addressNodeGot)
    }
}

func TestReachElementID_ElementIDNotExists_MustReturnFalse(t *testing.T) {
    element := `<div>I have no attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    docCNode := CNode{doc}
    
    addressNodeGot := docCNode.reachElementID("content")

    if addressNodeGot != nil {
        t.Errorf("Null pointer address expected.")
    }
}

func TestGetElementByID_ElementIDNotExists_MustReturnError(t *testing.T) {
    element := `<div>I have no attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    docCNode := CNode{doc}
    
    _, err := docCNode.GetElementByID("content")

    if err == nil {
        t.Errorf("Expected an error.")
    }
}

func TestGetElementByID_ElementIDExists_MustReturnSamePointerAddress(t *testing.T) {
    element := `<div id="content">I have attribute ID.</div>`
    
    doc, _ := html.Parse(strings.NewReader(element))
    
    divElementNode := doc.FirstChild.LastChild.FirstChild
    
    if divElementNode.Data != "div" {
        t.Fatal("Should be a div element.")
    }
    
    docCNode := CNode{doc}
    
    addressNodeExpected := divElementNode
    addressNodeGot, err := docCNode.GetElementByID("content")
    
    if err != nil {
        t.Fatal("Unexpected error.")
    }

    if addressNodeExpected != addressNodeGot.Node {
        t.Errorf("Pointer address expected [%p] but got [%p].", addressNodeExpected, addressNodeGot)
    }
}
