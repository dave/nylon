// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates parsing HTML data and walking the resulting tree.
package nml_test

import (
	"fmt"
	"log"
	"strings"

	"nml"
)

func ExampleParse() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(nml.Node)
	f = func(n nml.Node) {
		if n.GetType() == html.ElementNode && n.GetData() == "a" {
			for _, a := range n.GetAttr() {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.GetFirstChild(); c != nil; c = c.GetNextSibling() {
			f(c)
		}
	}
	f(doc)
	// Output:
	// foo
	// /bar/baz
}
