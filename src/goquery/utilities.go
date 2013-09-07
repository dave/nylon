package goquery

import (
	"nml"
)

func getChildren(n nml.Node) (result []nml.Node) {
	for c := n.GetFirstChild(); c != nil; c = c.GetNextSibling() {
		result = append(result, c)
	}
	return
}

// Loop through all container nodes to search for the target node.
func sliceContains(container []nml.Node, contained nml.Node) bool {
	for _, n := range container {
		if nodeContains(n, contained) {
			return true
		}
	}

	return false
}

// Checks if the contained node is within the container node.
func nodeContains(container nml.Node, contained nml.Node) bool {
	// Check if the parent of the contained node is the container node, traversing
	// upward until the top is reached, or the container is found.
	for contained = contained.GetParent(); contained != nil; contained = contained.GetParent() {
		if container == contained {
			return true
		}
	}
	return false
}

// Checks if the target node is in the slice of nodes.
func isInSlice(slice []nml.Node, node nml.Node) bool {
	return indexInSlice(slice, node) > -1
}

// Returns the index of the target node in the slice, or -1.
func indexInSlice(slice []nml.Node, node nml.Node) int {
	if node != nil {
		for i, n := range slice {
			if n == node {
				return i
			}
		}
	}
	return -1
}

// Appends the new nodes to the target slice, making sure no duplicate is added.
// There is no check to the original state of the target slice, so it may still
// contain duplicates. The target slice is returned because append() may create
// a new underlying array.
func appendWithoutDuplicates(target []nml.Node,
	nodes []nml.Node) []nml.Node {

	for _, n := range nodes {
		if !isInSlice(target, n) {
			target = append(target, n)
		}
	}

	return target
}

// Loop through a selection, returning only those nodes that pass the predicate
// function.
func grep(sel *Selection, predicate func(i int, s *Selection) bool) (result []nml.Node) {

	for i, n := range sel.Nodes {
		if predicate(i, newSingleSelection(n, sel.document)) {
			result = append(result, n)
		}
	}
	return
}

// Creates a new Selection object based on the specified nodes, and keeps the
// source Selection object on the stack (linked list).
func pushStack(fromSel *Selection, nodes []nml.Node) (result *Selection) {
	result = &Selection{nodes, fromSel.document, fromSel}
	return
}
