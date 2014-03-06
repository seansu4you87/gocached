package server

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Context struct {
	URL   string
	Count int
}

type IndexHandler struct {
	sync.Mutex
	count int
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int

	h.Lock()
	h.count++
	count = h.count
	h.Unlock()

	t, err := template.ParseFiles("server/index.html")
	if err != nil {
		fmt.Fprintf(w, "SERVER ERROR")
		return
	}
	c := &Context{URL: r.URL.Path[1:], Count: count}
	t.Execute(w, c)
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("index.html")
  c := &Context{URL: r.URL.Path[1:]}
  t.Execute(w, c)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":1988", nil)
}
*/
