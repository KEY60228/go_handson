package main

import (
	"bufio"
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

func input(s string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(s, ":")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	dsn := fmt.Sprintf("host=pgsql dbname=go_handson user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	nm := input("name")
	ml := input("mail")
	age, _ := strconv.Atoi(input("age"))

	q := "INSERT INTO mydata (name, mail, age) VALUES ($1, $2, $3)"

	conn.Exec(q, nm, ml, age)
	showRecord(conn)
}

func showRecord(conn *sql.DB) {
	q := "SELECT * FROM mydata"
	res, _ := conn.Query(q)
	for res.Next() {
		fmt.Println(mydatafmrws(res).Str())
	}
}

func mydatafmrws(res *sql.Rows) *Mydata {
	var md Mydata
	err := res.Scan(&md.Id, &md.Name, &md.Mail, &md.Age)
	if err != nil {
		panic(err)
	}
	return &md
}
