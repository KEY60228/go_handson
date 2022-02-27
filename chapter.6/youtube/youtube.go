package youtube

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = fmt.Sprintf("host=pgsql dbname=go_handson user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
var sesName = "ytboard-session"
var cs = sessions.NewCookieStore([]byte("secret-key-1234"))

func checkLogin(w http.ResponseWriter, r *http.Request) *User {
	ses, _ := cs.Get(r, sesName)
	if ses.Values["login"] == nil || !ses.Values["login"].(bool) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	ac := ""
	if ses.Values["account"] != nil {
		ac = ses.Values["account"].(string)
	}

	var user User
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.Where("account = $1", ac).First(&user)

	return &user
}

func notemp() *template.Template {
	tmp, _ := template.New("index").Parse("NO PAGE.")
	return tmp
}

func page(fname string) *template.Template {
	tmps, err := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	if err != nil {
		return notemp()
	}
	return tmps
}

func index(w http.ResponseWriter, r *http.Request) {
	user := checkLogin(w, r)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	var pl []Post
	db.Order("created_at desc").Limit(10).Find(&pl)
	var gl []Group
	db.Order("created_at desc").Limit(10).Find(&gl)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Plist   []Post
		Glist   []Group
	}{
		Title:   "Index",
		Message: "This is Top page.",
		Name:    user.Name,
		Account: user.Account,
		Plist:   pl,
		Glist:   gl,
	}
	err = page("index").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	user := checkLogin(w, r)

	pid := r.FormValue("pid")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		msg := r.PostFormValue("message")
		pId, _ := strconv.Atoi(pid)
		cmt := Comment{
			UserId:  int(user.Model.ID),
			PostId:  pId,
			Message: msg,
		}
		db.Create(&cmt)
	}

	var post Post
	var cmts []CommentJoin

	db.Where("id = $1", pid).First(&post)
	db.Table("comments").Select("comments.*, users.id, users.name").Joins("JOIN users ON users.id = comments.user_id").Where("comments.post_id = $1", pid).Order("created_at desc").Find(&cmts)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Post    Post
		Clist   []CommentJoin
	}{
		Title:   "Post",
		Message: "Post id = " + pid,
		Name:    user.Name,
		Account: user.Account,
		Post:    post,
		Clist:   cmts,
	}
	err = page("post").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	user := checkLogin(w, r)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		switch r.PostFormValue("form") {
		case "post":
			ad := r.PostFormValue("address")
			ad = strings.TrimSpace(ad)
			ad = strings.TrimPrefix(ad, "https://youtu.be/")

			pt := Post{
				UserId:  int(user.Model.ID),
				Address: ad,
				Message: r.PostFormValue("message"),
			}
			db.Create(&pt)
		case "group":
			gp := Group{
				UserId:  int(user.Model.ID),
				Name:    r.PostFormValue("name"),
				Message: r.PostFormValue("message"),
			}
			db.Create(&gp)
		}
	}

	var pts []Post
	var gps []Group
	db.Where("user_id = $1", user.ID).Order("created_at desc").Limit(10).Find(&pts)
	db.Where("user_id = $1", user.ID).Order("created_at desc").Limit(10).Find(&gps)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Plist   []Post
		Glist   []Group
	}{
		Title:   "Home",
		Message: "User account\"" + user.Account + "\".",
		Name:    user.Name,
		Account: user.Account,
		Plist:   pts,
		Glist:   gps,
	}
	err = page("home").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func group(w http.ResponseWriter, r *http.Request) {
	user := checkLogin(w, r)

	gid := r.FormValue("gid")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		ad := r.PostFormValue("address")
		ad = strings.TrimSpace(ad)
		ad = strings.TrimPrefix(ad, "https://youtu.be/")

		gId, _ := strconv.Atoi(gid)
		post := Post{
			UserId:  int(user.Model.ID),
			Address: ad,
			Message: r.PostFormValue("message"),
			GroupId: gId,
		}
		db.Create(&post)
	}

	var group Group
	var posts []Post
	db.Where("id = $1", gid).First(&group)
	db.Order("created_at desc").Where("group_id = $1", gid).Find(&posts)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Group   Group
		Plist   []Post
	}{
		Title:   "Group",
		Message: "Group id = " + gid,
		Name:    user.Name,
		Account: user.Account,
		Group:   group,
		Plist:   posts,
	}
	err = page("group").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	item := struct {
		Title   string
		Message string
		Account string
	}{
		Title:   "Login",
		Message: "type your account & password: ",
		Account: "",
	}

	if r.Method == "GET" {
		err := page("login").Execute(w, item)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if r.Method == "POST" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatal(err)
		}

		account := r.PostFormValue("account")
		pass := r.PostFormValue("password")
		item.Account = account

		var re int64
		var user User
		db.Where("account = $1 AND password = $2", account, pass).Find(&user).Count(&re)

		if re <= 0 {
			item.Message = "Wrong account or password."
			page("login").Execute(w, item)
			return
		}

		ses, _ := cs.Get(r, sesName)
		ses.Values["login"] = true
		ses.Values["account"] = account
		ses.Values["name"] = user.Name
		ses.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	err := page("login").Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	ses, _ := cs.Get(r, sesName)
	ses.Values["login"] = nil
	ses.Values["account"] = nil
	ses.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func Run() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/post", post)
	http.HandleFunc("/group", group)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":50510", nil)
}
