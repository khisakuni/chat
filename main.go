package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP requet
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var host = flag.String("host", ":8080", "The host of the application.")
	flag.Parse()
	r := newRoom()
	// r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)
	go r.run()
	log.Println("Starting web server on", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
