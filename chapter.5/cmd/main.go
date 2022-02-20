package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mydata struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Tel  string `json:"tel"`
}

func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + ", " + m.Tel + ">"
}

func main() {
	res, err := http.Get("https://tuyano-dummy-data.firebaseio.com/mydata.json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var items []Mydata

	err = json.Unmarshal(s, &items)
	if err != nil {
		panic(err)
	}

	for i, im := range items {
		fmt.Println(i, im.Str())
	}
}
