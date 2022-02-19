package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rt := func(f *os.File) {
		r := bufio.NewReaderSize(f, 4096)
		for i := 1; true; i++ {
			s, _, err := r.ReadLine()
			if err != nil {
				break
			}
			fmt.Println(i, ": ", string(s))
		}
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
