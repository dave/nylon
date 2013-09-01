// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nml

import (
	"nml/atom"
)

// A NodeType is the type of a Node.
type NodeType uint32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
	scopeMarkerNode
)

type Node interface {
	Parent() Node
	FirstChild() Node
	LastChild() Node
	PrevSibling() Node
	NextSibling() Node

	SetParent(el Node)
	SetFirstChild(el Node)
	SetLastChild(el Node)
	SetPrevSibling(el Node)
	SetNextSibling(el Node)

	Type() NodeType
	DataAtom() atom.Atom
	Data() string
	Namespace() string
	Attr() []Attribute

	SetType(nodeType NodeType)
	SetDataAtom(dataAtom atom.Atom)
	SetData(data string)
	SetNamespace(namespace string)
	SetAttr(attr []Attribute)

	clone(lookup func(node *NodeStruct) Node) Node

	BeforeRender()
}


// Section 12.2.3.3 says "scope markers are inserted when entering applet
// elements, buttons, object elements, marquees, table cells, and table
// captions, and are used to prevent formatting from 'leaking'".
var scopeMarker = NodeStruct{nodeType: scopeMarkerNode}

// A Node consists of a NodeType and some Data (tag name for element nodes,
// content for text) and are part of a tree of Nodes. Element nodes may also
// have a Namespace and contain a slice of Attributes. Data is unescaped, so
// that it looks like "a<b" rather than "a&lt;b". For element nodes, DataAtom
// is the atom for Data, or zero if Data is not a known tag name.
//
// An empty Namespace implies a "http://www.w3.org/1999/xhtml" namespace.
// Similarly, "math" is short for "http://www.w3.org/1998/Math/MathML", and
// "svg" is short for "http://www.w3.org/2000/svg".
type NodeStruct struct {
	parent, firstChild, lastChild, prevSibling, nextSibling Node

	nodeType  NodeType
	dataAtom  atom.Atom
	data      string
	namespace string
	attr      []Attribute
}
func (n *NodeStruct) Parent() Node {return n.parent}
func (n *NodeStruct) FirstChild() Node {return n.firstChild}
func (n *NodeStruct) LastChild() Node {return n.lastChild}
func (n *NodeStruct) PrevSibling() Node {return n.prevSibling}
func (n *NodeStruct) NextSibling() Node {return n.nextSibling}
func (n *NodeStruct) SetParent(el Node) {n.parent = el}
func (n *NodeStruct) SetFirstChild(el Node) {n.firstChild = el}
func (n *NodeStruct) SetLastChild(el Node) {n.lastChild = el}
func (n *NodeStruct) SetPrevSibling(el Node) {n.prevSibling = el}
func (n *NodeStruct) SetNextSibling(el Node) {n.nextSibling = el}
func (n *NodeStruct) Type() NodeType {return n.nodeType}
func (n *NodeStruct) DataAtom() atom.Atom {return n.dataAtom}
func (n *NodeStruct) Data() string {return n.data}
func (n *NodeStruct) Namespace() string {return n.namespace}
func (n *NodeStruct) Attr() []Attribute {return n.attr}
func (n *NodeStruct) SetType(nodeType NodeType) {n.nodeType = nodeType}
func (n *NodeStruct) SetDataAtom(dataAtom atom.Atom) {n.dataAtom = dataAtom}
func (n *NodeStruct) SetData(data string) {n.data = data}
func (n *NodeStruct) SetNamespace(namespace string) {n.namespace = namespace}
func (n *NodeStruct) SetAttr(attr []Attribute) {n.attr = attr}

// InsertBefore inserts newChild as a child of n, immediately before oldChild
// in the sequence of n's children. oldChild may be nil, in which case newChild
// is appended to the end of n's children.
//
// It will panic if newChild already has a parent or siblings.
func InsertBefore(node, newChild, oldChild Node) {
	if newChild.Parent() != nil || newChild.PrevSibling() != nil || newChild.NextSibling() != nil {
		panic("html: InsertBefore called for an attached child Node")
	}
	var prev, next Node
	if oldChild != nil {
		prev, next = oldChild.PrevSibling(), oldChild
	} else {
		prev = node.LastChild()
	}
	if prev != nil {
		prev.SetNextSibling(newChild)
	} else {
		node.SetFirstChild(newChild)
	}
	if next != nil {
		next.SetPrevSibling(newChild)
	} else {
		node.SetLastChild(newChild)
	}
	newChild.SetParent(node)
	newChild.SetPrevSibling(prev)
	newChild.SetNextSibling(next)
}

// AppendChild adds a node c as a child of n.
//
// It will panic if c already has a parent or siblings.
func AppendChild(node, child Node) {
	if child.Parent() != nil || child.PrevSibling() != nil || child.NextSibling() != nil {
		panic("html: AppendChild called for an attached child Node")
	}
	last := node.LastChild()
	if last != nil {
		last.SetNextSibling(child)
	} else {
		node.SetFirstChild(child)
	}
	node.SetLastChild(child)
	child.SetParent(node)
	child.SetPrevSibling(last)
}

// RemoveChild removes a node c that is a child of n. Afterwards, c will have
// no parent and no siblings.
//
// It will panic if c's parent is not n.
func RemoveChild(node, child Node) {
	if child.Parent() != node {
		panic("html: RemoveChild called for a non-child Node")
	}
	if node.FirstChild() == child {
		node.SetFirstChild(child.NextSibling())
	}
	if child.NextSibling() != nil {
		child.NextSibling().SetPrevSibling(child.PrevSibling())
	}
	if node.LastChild() == child {
		node.SetLastChild(child.PrevSibling())
	}
	if child.PrevSibling() != nil {
		child.PrevSibling().SetNextSibling(child.NextSibling())
	}
	child.SetParent(nil)
	child.SetPrevSibling(nil)
	child.SetNextSibling(nil)
}

// reparentChildren reparents all of src's child nodes to dst.
func reparentChildren(dst, src Node) {
	for {
		child := src.FirstChild()
		if child == nil {
			break
		}
		RemoveChild(src, child)
		AppendChild(dst, child)
	}
}

// clone returns a new node with the same type, data and attributes.
// The clone has no parent, no siblings and no children.
func (n *NodeStruct) clone(lookup func(node *NodeStruct) Node) Node {

	attr := make([]Attribute, len(n.Attr()))
	copy(attr, n.Attr())

	m := lookup(&NodeStruct{
		nodeType: n.Type(),
		dataAtom: n.DataAtom(),
		data:     n.Data(),
		attr:     attr,
	})
	return m
}

// nodeStack is a stack of nodes.
type nodeStack []Node

// pop pops the stack. It will panic if s is empty.
func (s *nodeStack) pop() Node {
	i := len(*s)
	n := (*s)[i-1]
	*s = (*s)[:i-1]
	return n
}

// top returns the most recently pushed node, or nil if s is empty.
func (s *nodeStack) top() Node {
	if i := len(*s); i > 0 {
		return (*s)[i-1]
	}
	return nil
}

// index returns the index of the top-most occurrence of n in the stack, or -1
// if n is not present.
func (s *nodeStack) index(n Node) int {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i] == n {
			return i
		}
	}
	return -1
}

// insert inserts a node at the given index.
func (s *nodeStack) insert(i int, n Node) {
	(*s) = append(*s, nil)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = n
}

// remove removes a node from the stack. It is a no-op if n is not present.
func (s *nodeStack) remove(n Node) {
	i := s.index(n)
	if i == -1 {
		return
	}
	copy((*s)[i:], (*s)[i+1:])
	j := len(*s) - 1
	(*s)[j] = nil
	*s = (*s)[:j]
}
