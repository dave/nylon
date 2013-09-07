package goquery

import (
	"nml"
	"net/http"
	"net/url"
)

// Document represents an HTML document to be manipulated. Unlike jQuery, which
// is loaded as part of a DOM document, and thus acts upon its containing
// document, GoQuery doesn't know which HTML document to act upon. So it needs
// to be told, and that's what the Document class is for. It holds the root
// document node to manipulate, and can make selections on this document.
type Document struct {
	*Selection
	Url      *url.URL
	rootNode nml.Node
}

// NewDocumentFromNode() is a Document constructor that takes a root nml Node
// as argument.
func NewDocumentFromNode(root nml.Node) (d *Document) {
	return newDocument(root, nil)
}

// NewDocument() is a Document constructor that takes a string URL as argument.
// It loads the specified document, parses it, and stores the root Document
// node, ready to be manipulated.
func NewDocument(url string, lookup func(node *nml.NodeStruct) nml.Node) (d *Document, e error) {
	// Load the URL
	res, e := http.Get(url)
	if e != nil {
		return
	}
	return NewDocumentFromResponse(res, lookup)
}

// NewDocumentFromResponse() is another Document constructor that takes an http resonse as argument.
// It loads the specified response's document, parses it, and stores the root Document
// node, ready to be manipulated.
func NewDocumentFromResponse(res *http.Response, lookup func(node *nml.NodeStruct) nml.Node) (d *Document, e error) {
	defer res.Body.Close()

	// Parse the HTML into nodes
	root, e := nml.Parse(res.Body, lookup)
	if e != nil {
		return
	}

	// Create and fill the document
	d = newDocument(root, res.Request.URL)
	return
}

// Private constructor, make sure all fields are correctly filled.
func newDocument(root nml.Node, url *url.URL) (d *Document) {
	// Create and fill the document
	d = &Document{nil, url, root}
	d.Selection = newSingleSelection(root, d)
	return
}

// Selection represents a collection of nodes matching some criteria. The
// initial Selection can be created by using Document.Find(), and then
// manipulated using the jQuery-like chainable syntax and methods.
type Selection struct {
	Nodes    []nml.Node
	document *Document
	prevSel  *Selection
}

// Helper constructor to create an empty selection
func newEmptySelection(doc *Document) *Selection {
	return &Selection{nil, doc, nil}
}

// Helper constructor to create a selection of only one node
func newSingleSelection(node nml.Node, doc *Document) *Selection {
	return &Selection{[]nml.Node{node}, doc, nil}
}
