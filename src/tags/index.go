// First we run PreParse on each element, starting from the deepest. If we need to
// alter the contents, we should do it here
// We run PreRender on each element, starting from the root.
package tags

import (
	"nml"
)

func Index(node *nml.NodeStruct) nml.Node {
	switch node.GetData() {
	case "my-tag":
		return &MyTag{NodeStruct: node}
	case "my-bio":
		return &MyBio{NodeStruct: node}
	}
	return &Tag{NodeStruct: node}
}

type Tag struct {
	*nml.NodeStruct
}
