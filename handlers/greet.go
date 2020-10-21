package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Greet struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewGreet(l *log.Logger) *Greet {
	return &Greet{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Greet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad request data", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s.\n", b)
}
