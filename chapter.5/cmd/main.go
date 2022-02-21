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

	id, _ := strconv.Atoi(input("update ID"))

	q := "SELECT * FROM mydata WHERE id = $1"

	row := conn.QueryRow(q, id)
	target := mydatafmrw(row)

	nm := input("name(" + target.Name + ")")
	ml := input("mail(" + target.Mail + ")")
	inputAge := input("age(" + strconv.Itoa(target.Age) + ")")

	age, _ := strconv.Atoi(inputAge)
	if nm == "" {
		nm = target.Name
	}
	if ml == "" {
		ml = target.Mail
	}
	if inputAge == "" {
		age = target.Age
	}

	q = "UPDATE mydata SET name = $1, mail = $2, age = $3 WHERE id = $4"
	conn.Exec(q, nm, ml, age, id)

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

func mydatafmrw(res *sql.Row) *Mydata {
	var md Mydata
	err := res.Scan(&md.Id, &md.Name, &md.Mail, &md.Age)
	if err != nil {
		panic(err)
	}
	return &md
}
