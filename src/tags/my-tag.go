package tags

import (
	"nml"
	"goquery"
	"errors"
)

type MyTag struct {
	*nml.NodeStruct
	Me *MyBio
}

func (n *MyTag) Init() error {
	err := n.NodeStruct.Init(); if err != nil { return err }
	g := goquery.NewDocumentFromNode(n)
	me, ok := g.Selection.Find("#Me").Nodes[0].(*MyBio); if !ok { return errors.New("#Me is not a MyBio") }
	n.Me = me
	n.Me.Color = "foo"
	n.Logger.Info("%#v", n.Me)
	return nil
}

func (n *MyTag) Render() {
	n.Logger.Info("%#v", n.Me)
	me := n.Me
	me.Color = "red"
}
