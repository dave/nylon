package root

import (
	"store"
	"net/http"
	"nml"
	"strings"
	"bufio"
	"tags"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	str := store.Get()
	read := strings.NewReader(str)

	doc, err := nml.Parse(read, tags.Index)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(w)

	err = nml.Render(buf, doc)
	buf.Flush()

}

