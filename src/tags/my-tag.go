package tags

import (
	"common"
	"code.google.com/p/go.net/html"
	"render"
)

type My_Tag struct {
	*html.Node
}
func (n *My_Tag) Render(w common.Writer, wrapper func(*html.Node) common.Renderer) error {

	if err := render.OpeningTag(w, n.Node); err != nil {
		return err
	}

	if err := render.Contents(w, n.Node, wrapper); err != nil {
		return err
	}

	w.WriteString(" BAR!!!")

	if err := render.ClosingTag(w, n.Node); err != nil {
		return err
	}

	return nil

}
