package util

import (
    _ "strings"
    "testing"
    _ "golang.org/x/net/html"
)

func TestExtractFolderMessage_HTMLPageHasATextAreaElement_MustReturnTheContentOfTextArea_MustReturnNoError(t *testing.T) {
    htmlPage := `<html><head></head><body><textarea id="text">Hello World.</textarea></body></html>`

    content, err := ExtractFolderMessage(htmlPage)

    if err != nil {
        t.Errorf("Should not return an error.\n")
    }

    expected := "Hello World."

    if content != expected {
        t.Errorf("Expected value [%s] but got [%s].\n", expected, content)
    }
}

func TestExtractFolderMessage_HTMLPageHasAEmptyTextAreaElement_MustReturnFalse(t *testing.T) {
   htmlPage := `<html><head></head><body><textarea id="text"></textarea></body></html>`
   
   content, err := ExtractFolderMessage(htmlPage)
   
   if err != nil {
       t.Errorf("Should not return an error.\n")
   }
   
   if content != "" {
       t.Errorf("Expected false value but got [%s].\n", content)
   }
}
