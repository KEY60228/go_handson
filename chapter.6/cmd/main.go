package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

type Temps struct {
	notemp *template.Template
	index  *template.Template
	hello  *template.Template
}

func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps
}

func index(w http.ResponseWriter, r *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Top page.",
	}

	err := page("index").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	data := []string{"One", "Two", "Three"}

	item := struct {
		Title string
		Data  []string
	}{
		Title: "Send values",
		Data:  data,
	}

	err := page("hello").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { index(w, r) })
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { hello(w, r) })
	http.ListenAndServe(":50510", nil)
}
