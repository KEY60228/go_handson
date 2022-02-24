package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	html := `<html><body>
	<h1>Hello</h1>
	<p>This is sample message.</p>
	</body></html>`

	tf, err := template.New("index").Parse(html)
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
