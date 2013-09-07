package root

import (
	"store"
	"net/http"
	"nml"
	"bufio"
	"tags"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	reader := store.Get("root")
	doc, err := nml.Parse(reader, tags.Index); if err != nil { panic(err) }
	buf := bufio.NewWriter(w)
	err = nml.Render(buf, doc)
	buf.Flush()

}

