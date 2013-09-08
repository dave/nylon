package common

import (
	"net/http"
	"appengine"
)

type Logger struct {
	*http.Request
}

func (l *Logger) Info(f string, s ...interface{}) {
	c := appengine.NewContext(l.Request)
	c.Infof(f, s)
}
