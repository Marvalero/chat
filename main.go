package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/Marvalero/chat/chat"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.getTempl().Execute(w, r)
}

// This guarantees that the template rendering will only be executed once, regardless of how many goroutines are calling ServeHTTP
func (t *templateHandler) getTempl() *template.Template {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	return t.templ
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	mainRoom := chat.NewRoom()
	go mainRoom.Run()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", mainRoom)

	log.Println("Starting web server on", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
