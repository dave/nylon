package tags

import (
	"nml"
	"goquery"
)

type My_Tag struct {
	*nml.NodeStruct
}

func (n My_Tag) BeforeRender() {
	g := goquery.NewDocumentFromNode(n)

	g.Selection.SetAttr("style", "color:blue;")

}
