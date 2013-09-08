package tags

import (
	"nml"
	"store"
	"goquery"
	"fmt"
)

type MyBio struct {
	*nml.NodeStruct
	Color string
}

func (n *MyBio) Init() error {
	reader := store.Get("bio")
	children, err := nml.ParseFragmentBody(reader, Index); if err != nil { return err }
	nml.AppendChildren(n, children)
	return nil
}

func (n *MyBio) Render() {
	g := goquery.NewDocumentFromNode(n)
	g.Selection.SetAttr("style", fmt.Sprint("color:", n.Color, ";"))
}
