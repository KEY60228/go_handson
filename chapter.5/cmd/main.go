package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	var data []interface{}

	err = json.Unmarshal(s, &data)
	if err != nil {
		panic(err)
	}

	for i, im := range data {
		m := im.(map[string]interface{})
		fmt.Println(i, m["name"].(string), m["mail"].(string), m["tel"].(string))
	}
}
