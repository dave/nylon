package tags

import (
	"nml"
)

type My_Tag struct {
	*nml.NodeStruct
}

func (n My_Tag) BeforeRender() {
	n.NodeStruct.Attr = append(n.NodeStruct.Attr, nml.Attribute{"", "style", "color:red;"});
}
