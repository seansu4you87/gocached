package main

import (
	"html/template"
	"net/http"
)

type Context struct {
	URL string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	c := &Context{URL: r.URL.Path[1:]}
	t.Execute(w, c)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1988", nil)
}
