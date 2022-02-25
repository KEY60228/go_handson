package main

import (
	"html/template"
	"log"
	"net/http"
)

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

func setupTemp() *Temps {
	temps := &Temps{}

	temps.notemp = notemp()

	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		index = temps.notemp
	}
	temps.index = index

	hello, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		hello = temps.notemp
	}
	temps.hello = hello

	return temps
}

func index(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	err := tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	var flg bool
	item := struct {
		Flg      bool
		Title    string
		Message  string
		JMessage string
	}{
		Flg:      flg,
		Title:    "Send values",
		Message:  "This is Sample message.",
		JMessage: "これはサンプルです",
	}

	err := tmp.Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	temps := setupTemp()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, temps.index)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r, temps.hello)
	})
	http.ListenAndServe(":50510", nil)
}
