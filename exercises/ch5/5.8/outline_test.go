package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestIDEquals(t *testing.T) {
	// htmlString := `<html><head/><body><div id=thediv><span class="foo">hello world</span></div></body></html>`
	htmlString := `<div id=thediv><span class="foo">hello world</span></div>`
	// doc, err := html.Parse(strings.NewReader(`<div id=thediv><span class="foo">hello world</span></div>`))
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		t.Errorf("Failed to parse HTML:", err)
	}

	expected := true

	// document ->
	//   html ->
	//     head
	//     body ->
	//       div
	actual := IDEquals(doc.FirstChild.FirstChild.NextSibling.FirstChild, "thediv")

	if expected != actual {
		t.Errorf("Expected:\n%v\n\nGot:\n%v", expected, actual)
	}
}
