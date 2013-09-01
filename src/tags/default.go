package tags

import (
	"nml"
)

type Default struct {
	*nml.NodeStruct
}

func (n *Default) BeforeRender() {

}

func Lookup(node *nml.NodeStruct) nml.Node {
	switch node.GetData() {
	case "my-tag":
		return &My_Tag{node}
	default:
		return &Default{node}
	}
}
