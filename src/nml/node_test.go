// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nml

import (
	"fmt"
)

// checkTreeConsistency checks that a node and its descendants are all
// consistent in their parent/child/sibling relationships.
func checkTreeConsistency(n Node) error {
	return checkTreeConsistency1(n, 0)
}

func checkTreeConsistency1(n Node, depth int) error {
	if depth == 1e4 {
		return fmt.Errorf("html: tree looks like it contains a cycle")
	}
	if err := checkNodeConsistency(n); err != nil {
		return err
	}
	for c := n.GetFirstChild(); c != nil; c = c.GetNextSibling() {
		if err := checkTreeConsistency1(c, depth+1); err != nil {
			return err
		}
	}
	return nil
}

// checkNodeConsistency checks that a node's parent/child/sibling relationships
// are consistent.
func checkNodeConsistency(n Node) error {
	if n == nil {
		return nil
	}

	nParent := 0
	for p := n.GetParent(); p != nil; p = p.GetParent() {
		nParent++
		if nParent == 1e4 {
			return fmt.Errorf("html: parent list looks like an infinite loop")
		}
	}

	nForward := 0
	for c := n.GetFirstChild(); c != nil; c = c.GetNextSibling() {
		nForward++
		if nForward == 1e6 {
			return fmt.Errorf("html: forward list of children looks like an infinite loop")
		}
		if c.GetParent() != n {
			return fmt.Errorf("html: inconsistent child/parent relationship")
		}
	}

	nBackward := 0
	for c := n.GetLastChild(); c != nil; c = c.GetPrevSibling() {
		nBackward++
		if nBackward == 1e6 {
			return fmt.Errorf("html: backward list of children looks like an infinite loop")
		}
		if c.GetParent() != n {
			return fmt.Errorf("html: inconsistent child/parent relationship")
		}
	}

	if n.GetParent() != nil {
		if n.GetParent() == n {
			return fmt.Errorf("html: inconsistent parent relationship")
		}
		if n.GetParent() == n.GetFirstChild() {
			return fmt.Errorf("html: inconsistent parent/first relationship")
		}
		if n.GetParent() == n.GetLastChild() {
			return fmt.Errorf("html: inconsistent parent/last relationship")
		}
		if n.GetParent() == n.GetPrevSibling() {
			return fmt.Errorf("html: inconsistent parent/prev relationship")
		}
		if n.GetParent() == n.GetNextSibling() {
			return fmt.Errorf("html: inconsistent parent/next relationship")
		}

		parentHasNAsAChild := false
		for c := n.GetParent().GetFirstChild(); c != nil; c = c.GetNextSibling() {
			if c == n {
				parentHasNAsAChild = true
				break
			}
		}
		if !parentHasNAsAChild {
			return fmt.Errorf("html: inconsistent parent/child relationship")
		}
	}

	if n.GetPrevSibling() != nil && n.GetPrevSibling().GetNextSibling() != n {
		return fmt.Errorf("html: inconsistent prev/next relationship")
	}
	if n.GetNextSibling() != nil && n.GetNextSibling().GetPrevSibling() != n {
		return fmt.Errorf("html: inconsistent next/prev relationship")
	}

	if (n.GetFirstChild() == nil) != (n.GetLastChild() == nil) {
		return fmt.Errorf("html: inconsistent first/last relationship")
	}
	if n.GetFirstChild() != nil && n.GetFirstChild() == n.GetLastChild() {
		// We have a sole child.
		if n.GetFirstChild().GetPrevSibling() != nil || n.GetFirstChild().GetNextSibling() != nil {
			return fmt.Errorf("html: inconsistent sole child's sibling relationship")
		}
	}

	seen := map[Node]bool{}

	var last Node
	for c := n.GetFirstChild(); c != nil; c = c.GetNextSibling() {
		if seen[c] {
			return fmt.Errorf("html: inconsistent repeated child")
		}
		seen[c] = true
		last = c
	}
	if last != n.GetLastChild() {
		return fmt.Errorf("html: inconsistent last relationship")
	}

	var first Node
	for c := n.GetLastChild(); c != nil; c = c.GetPrevSibling() {
		if !seen[c] {
			return fmt.Errorf("html: inconsistent missing child")
		}
		delete(seen, c)
		first = c
	}
	if first != n.GetFirstChild() {
		return fmt.Errorf("html: inconsistent first relationship")
	}

	if len(seen) != 0 {
		return fmt.Errorf("html: inconsistent forwards/backwards child list")
	}

	return nil
}
