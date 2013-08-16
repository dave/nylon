package common

import (
	"io"
	"code.google.com/p/go.net/html"
)

type Renderer interface {
	Render(w Writer, wrapper func(*html.Node) Renderer) error
}

type Writer interface {
	io.Writer
	WriteByte(c byte) error // in Go 1.1, use io.ByteWriter
	WriteString(string) (int, error)
}
