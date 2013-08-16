package hello

import (
	"fmt"
	"foo"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := foo.Get()
	fmt.Fprint(w, "Hello, world!", s)
}
