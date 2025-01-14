package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	Name string
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL is %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "k:%q, v:%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "name %q\n", engine.Name)
		fmt.Fprintf(w, "404 not found: %q\n", r.URL)
		engine.Name = "not"
		fmt.Fprintf(w, "change name %q\n", engine.Name)
	}
}

func main() {
	var engine *Engine
	engine = new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
