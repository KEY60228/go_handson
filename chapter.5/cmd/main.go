package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Mydata struct {
	Id   int
	Name string
	Mail string
	Age  int
}

func (m *Mydata) Str() string {
	return "<" + strconv.Itoa(m.Id) + ": \"" + m.Name + "\" " + m.Mail + ", " + strconv.Itoa(m.Age) + ">"
}

func main() {
	dsn := fmt.Sprintf("host=pgsql dbname=go_handson user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	q := "SELECT * FROM mydata"
	res, err := conn.Query(q)
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var md Mydata
		err := res.Scan(&md.Id, &md.Name, &md.Mail, &md.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(md.Str())
	}
}
