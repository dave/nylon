package root

import (
	"store"
	"net/http"
	"code.google.com/p/go.net/html"
	"strings"
	"bufio"
	"render"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	str := store.Get()
	read := strings.NewReader(str)

	doc, err := html.Parse(read)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(w)

	node := WrapNode(doc)

	err = node.Render(buf, WrapNode)
	if err != nil && err != render.PlaintextAbort{
		panic(err)
	}
	buf.Flush()

}

