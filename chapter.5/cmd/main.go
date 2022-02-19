package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	p := "https://golang.org"
	re, err := http.Get(p)
	if err != nil {
		panic(err)
	}
	defer re.Body.Close()

	s, err := ioutil.ReadAll(re.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))
}
