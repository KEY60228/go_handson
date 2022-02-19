package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wt := func(f *os.File, s string) {
		_, err := f.WriteString(s + "\n")
		if err != nil {
			panic(err)
		}
	}
	fn := "data.txt"

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("*** start ***")
	wt(f, "*** start ***")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("type message: ")
		scanner.Scan()
		s := scanner.Text()

		if s == "" {
			break
		}
		wt(f, s)
	}
	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
