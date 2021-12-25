package example

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// type Node struct {
// 	Type                     NodeType
// 	Data                     string
// 	Attr                     []Attribute
// 	FirsthChild, NextSibling *Node
// }

// type NodeType int32

// const (
// 	ErrorNode NodeType = iota
// 	TextNode
// 	DocumentNode
// 	ElementNode
// 	CommentNode
// 	DoctypeNode
// )

// type Attribute struct {
// 	Key, Val string
// }

// func Parse(r io.Reader) (*Node, error)

func StudHTML() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
