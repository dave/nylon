package goquery

import (
	"bytes"
	"nml"
)

// Attr() gets the specified attribute's value for the first element in the
// Selection. To get the value for each element individually, use a looping
// construct such as Each() or Map() method.
func (this *Selection) Attr(attrName string) (val string, exists bool) {
	if len(this.Nodes) == 0 {
		return
	}
	return getAttributeValue(attrName, this.Nodes[0])
}

// SetAttr() sets the specified attribute's value for all elements in the
// Selection.
func (this *Selection) SetAttr(attrName string, attrValue string) {
	if len(this.Nodes) == 0 {
		return
	}
	for _, node := range this.Nodes {
		setAttributeValue(attrName, attrValue, node)
	}
}

// Text() gets the combined text contents of each element in the set of matched
// elements, including their descendants.
func (this *Selection) Text() string {
	var buf bytes.Buffer

	// Slightly optimized vs calling Each(): no single selection object created
	for _, n := range this.Nodes {
		buf.WriteString(getNodeText(n))
	}
	return buf.String()
}

// Size() is an alias for Length().
func (this *Selection) Size() int {
	return this.Length()
}

// Length() returns the number of elements in the Selection object.
func (this *Selection) Length() int {
	return len(this.Nodes)
}

// Html() gets the HTML contents of the first element in the set of matched
// elements. It includes text and comment nodes.
func (this *Selection) Html() (ret string, e error) {
	// Since there is no .innerHtml, the HTML content must be re-created from
	// the nodes usint nml.Render().
	var buf bytes.Buffer

	if len(this.Nodes) > 0 {
		for c := this.Nodes[0].GetFirstChild(); c != nil; c = c.GetNextSibling() {
			e = nml.Render(&buf, c)
			if e != nil {
				return
			}
		}
		ret = buf.String()
	}

	return
}

// Get the specified node's text content.
func getNodeText(node nml.Node) string {
	if node.GetType() == nml.TextNode {
		// Keep newlines and spaces, like jQuery
		return node.GetData()
	} else if node.GetFirstChild() != nil {
		var buf bytes.Buffer
		for c := node.GetFirstChild(); c != nil; c = c.GetNextSibling() {
			buf.WriteString(getNodeText(c))
		}
		return buf.String()
	}

	return ""
}

// Private function to get the specified attribute's value from a node.
func getAttributeValue(attrName string, n nml.Node) (val string, exists bool) {
	if n == nil {
		return
	}

	for _, a := range n.GetAttr() {
		if a.Key == attrName {
			val = a.Val
			exists = true
			return
		}
	}
	return
}

// Private function to get the specified attribute's value from a node.
func setAttributeValue(attrName string, attrValue string, n nml.Node) {
	if n == nil {
		return
	}

	attr := n.GetAttr()

	for _, a := range attr {
		if a.Key == attrName {
			a.Val = attrValue
			n.SetAttr(attr)
			return
		}
	}
	n.SetAttr(append(attr, nml.Attribute{Key: attrName, Val: attrValue}))
	return
}
