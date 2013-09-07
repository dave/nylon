package tags

import (
	"nml"
)

func Index(node *nml.NodeStruct) nml.Node {
	switch node.GetData() {
	case "my-tag":
		return &My_Tag{node}
	default:
		return &Raw{node}
	}
}
