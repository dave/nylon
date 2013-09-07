package tags

import (
	"nml"
	"store"
)

type MyBio struct {
	*nml.NodeStruct
}

func (n MyBio) Init() error {
	reader := store.Get("bio")
	children, err := nml.ParseFragmentBody(reader, Index); if err != nil { return err }
	nml.AppendChildren(n, children)
	return nil
}

func (n MyBio) Render() { }
