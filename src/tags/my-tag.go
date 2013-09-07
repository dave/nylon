package tags

import (
	"nml"
	"goquery"
	//"errors"
)

type MyTag struct {
	*nml.NodeStruct
	Me nml.Node
}

func (n MyTag) Init() error {
	err := n.NodeStruct.Init(); if err != nil { return err }
	/*
	g := goquery.NewDocumentFromNode(n)
	me, ok := g.Find("#Me").Nodes[0].(MyBio)
	if !ok {
		return errors.New("#Me is not a MyBio")
	}
	n.Me = me
	*/
	g := goquery.NewDocumentFromNode(n)
	n.Me = g.Find("#Me").Nodes[0]
	return nil
}

func (n MyTag) Render() {
	g := goquery.NewDocumentFromNode(n)
	g.Find("p").SetAttr("style", "color:green;")
}
