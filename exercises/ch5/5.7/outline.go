// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

// func main() {
//   for _, url := range os.Args[1:] {
//     fmt.Println(PrettyPrint(url))
//   }
// }

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// not thread safe
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	// trimmed := bytes.Trim(buf.Bytes(), "\x00")
	// trimmed = bytes.TrimSpace(trimmed)
	// return string(trimmed)

	return string(strings.TrimSpace(buf.String()))
}

func TidyHTML(doc *html.Node) string {
	out := captureStdout(func() {
		forEachNode(doc, startElement, endElement)
	})

	return out
}

func TidyHTMLString(s string) (string, error) {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return "", err
	}

	return TidyHTML(doc), nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		var attrs string

		for _, a := range n.Attr {
			attrs += " " + a.Key + "='" + a.Val + "'"
		}

		if n.FirstChild != nil {
			// opening tag:
			//  print <a> and increase depth
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrs)
			depth++
		} else {
			// self-closing tag:
			//   print <img/> and don't increase depth
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrs)
		}
	case html.TextNode:
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!--\n", depth*2, "")
		depth++

		cleanText := strings.Replace(strings.TrimSpace(n.Data), "\n", fmt.Sprintf("\n%*s", depth*2, ""), -1)
		fmt.Printf("%*s%s\n", depth*2, "", cleanText)
		depth++
	default:
		return
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	case html.TextNode:
		depth--
		cleanText := strings.Replace(strings.TrimSpace(n.Data), "\n", fmt.Sprintf("\n%*s", depth*2, ""), -1)
		fmt.Printf("%*s%s\n", depth*2, "", cleanText)

	case html.CommentNode:
		depth--
		depth--
		fmt.Printf("%*s-->\n", depth*2, "")

	}
}

//!-startend
