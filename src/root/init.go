package root

import (
	"store"
	"net/http"
	"nml"
	"bufio"
	"tags"
	"common"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	logger := &common.Logger{r}

	reader := store.Get("root")
	doc, err := nml.Parse(reader, tags.Index, logger); if err != nil { panic(err) }
	buf := bufio.NewWriter(w)
	err = nml.Render(buf, doc)
	buf.Flush()

}

