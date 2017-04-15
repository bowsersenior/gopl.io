package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func printText(n *html.Node) {
	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)

		if s != "" {
			fmt.Println(s)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			// don't descend into script or style elms
			continue
		}

		printText(c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	printText(doc)
}
