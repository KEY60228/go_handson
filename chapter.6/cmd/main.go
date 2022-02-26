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
	msg := "login name and password: "

	ses, _ := cs.Get(r, "hello-session")

	if r.Method == "POST" {
		ses.Values["login"] = nil
		ses.Values["name"] = nil

		name := r.PostFormValue("name")
		pw := r.PostFormValue("password")
		if name == pw {
			ses.Values["login"] = true
			ses.Values["name"] = name
		}
		ses.Save(r, w)
	}

	flg, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)
	if flg {
		msg = "logined: " + lname
	}

	item := struct {
		Title   string
		Message string
	}{
		Title:   "Send values",
		Message: msg,
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
