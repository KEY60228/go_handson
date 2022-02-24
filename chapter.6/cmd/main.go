package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tf, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		log.Fatal(err)
	}

	hh := func(w http.ResponseWriter, r *http.Request) {
		err := tf.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/hello", hh)
	http.ListenAndServe(":50510", nil)
}
