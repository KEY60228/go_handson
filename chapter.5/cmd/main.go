package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://golang.org")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("a").Each(func(n int, sel *goquery.Selection) {
		lk, _ := sel.Attr("href")
		fmt.Println(n, sel.Text(), "(", lk, ")")
	})
}
