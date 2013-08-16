package tags

import (
	"common"
	"code.google.com/p/go.net/html"
	"render"
)

type Default struct {
	*html.Node
}
func (n *Default) Render(w common.Writer, wrapper func(*html.Node) common.Renderer) error {

	if n.Type == html.ElementNode || n.Type == html.DocumentNode {

		if err := render.OpeningTag(w, n.Node); err != nil {
			return err
		}

		if err := render.Contents(w, n.Node, wrapper); err != nil {
			return err
		}

		if err := render.ClosingTag(w, n.Node); err != nil {
			return err
		}

	} else {

		if err := render.NonElementNode(w, n.Node); err != nil {
			return err
		}

	}
	return nil

}
