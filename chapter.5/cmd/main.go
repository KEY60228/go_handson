package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rt := func(f *os.File) {
		s, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(s))
	}

	fn := "data.txt"

	f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("<<< start >>>")
	rt(f)
	fmt.Println("<<< end >>>")
}
