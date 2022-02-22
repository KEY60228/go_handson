package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
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
	a := app.New()
	w := a.NewWindow("app")
	a.Settings().SetTheme(theme.DarkTheme())

	edit := widget.NewMultiLineEntry()
	sc := widget.NewScrollContainer(edit)
	find := widget.NewEntry()
	info := widget.NewLabel("infomation bar.")

	showInfo := func(s string) {
		info.SetText(s)
		dialog.ShowInformation("info", s, w)
	}

	checkErr := func(err error) bool {
		if err != nil {
			info.SetText(err.Error())
			return true
		}
		return false
	}

	setDB := func() *sql.DB {
		dsn := fmt.Sprintf("host=pgsql dbname=go_handson user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
		conn, err := sql.Open("postgres", dsn)
		if checkErr(err) {
			return nil
		}
		return conn
	}

	nf := func() {
		dialog.ShowConfirm("Alert", "Clear form?", func(f bool) {
			if f {
				find.SetText("")
				w.SetTitle("App")
				edit.SetText("")
			}
		}, w)
	}

	wf := func() {
		fstr := find.Text
		if !strings.HasPrefix(fstr, "http") {
			fstr = "http://" + fstr
			find.SetText(fstr)
		}

		dc, err := goquery.NewDocument(fstr)
		if checkErr(err) {
			return
		}

		title := dc.Find("title")
		w.SetTitle(title.Text())

		html, err := dc.Html()
		if checkErr(err) {
			return
		}

		cvtr := md.NewConverter("", true, nil)
		mkdn, err := cvtr.ConvertString(html)
		if checkErr(err) {
			return
		}

		edit.SetText(mkdn)
		info.SetText("get web data.")
	}

	ff := func() {
		q := "SELECT * FROM md_data WHERE title LIKE %$1%"
		conn := setDB()
		if conn == nil {
			return
		}
		defer conn.Close()

		res, err := conn.Query(q, find.Text)
		if checkErr(err) {
			return
		}

		rs := ""
		for res.Next() {
			var ID int
			var TT string
			var UR string
			var MR string
			err := res.Scan(&ID, &TT, &UR, &MR)
			if checkErr(err) {
				return
			}
			rs += strconv.Itoa(ID) + ":" + TT + "\n"
		}
		edit.SetText(rs)
		info.SetText("Find: " + find.Text)
	}

	idf := func(id int) {
		q := "SELECT * FROM md_data WHERE id = $1"
		conn := setDB()
		if conn == nil {
			return
		}
		defer conn.Close()

		res := conn.QueryRow(q, id)

		var ID int
		var TT string
		var UR string
		var MR string
		res.Scan(&ID, &TT, &UR, &MR)

		w.SetTitle(TT)
		find.SetText(UR)
		edit.SetText(MR)
		info.SetText("Find id = " + strconv.Itoa(ID) + ".")
	}

	sf := func() {
		dialog.ShowConfirm("Alert", "Save data?", func(f bool) {
			if f {
				conn := setDB()
				if conn == nil {
					return
				}
				defer conn.Close()

				q := "INSERT INTO md_data (title, url, markdonw) VALUES ($1, $2, $3)"
				_, err := conn.Exec(q, w.Title(), find.Text, edit.Text)
				if checkErr(err) {
					return
				}
				showInfo("Save data to database!")
			}
		}, w)
	}

	xf := func() {
		dialog.ShowConfirm("Alert", "Export this data?", func(f bool) {
			if f {
				fn := w.Title() + ".md"
				ctt := "# " + w.Title() + "\n\n"
				ctt += "## " + find.Text + "\n\n"
				ctt += edit.Text
				err := ioutil.WriteFile(fn, []byte(ctt), os.ModePerm)
				if checkErr(err) {
					return
				}
				showInfo("Export data to file \"" + fn + "\".")
			}
		}, w)
	}

	qf := func() {
		dialog.ShowConfirm("Alert", "Quit application?", func(f bool) {
			if f {
				a.Quit()
			}
		}, w)
	}

	tf := true

	cf := func() {
		if tf {
			a.Settings().SetTheme(theme.LightTheme())
			info.SetText("change to Light-Theme.")
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			info.SetText("change to Dark-Theme.")
		}
		tf = !tf
	}

	cbtn := widget.NewButton("Clear", func() { nf() })
	wbtn := widget.NewButton("Get Web", func() { wf() })
	fbtn := widget.NewButton("Find data", func() { ff() })
	ibtn := widget.NewButton("Get ID data", func() {
		rid, err := strconv.Atoi(find.Text)
		if checkErr(err) {
			return
		}
		idf(rid)
	})
	sbtn := widget.NewButton("Save data", func() { sf() })
	xbtn := widget.NewButton("Export data", func() { xf() })

	createMenubar := func() *fyne.MainMenu {
		return fyne.NewMainMenu(
			fyne.NewMenu(
				"File",
				fyne.NewMenuItem("New", func() { nf() }),
				fyne.NewMenuItem("Get Web", func() { wf() }),
				fyne.NewMenuItem("Find", func() { ff() }),
				fyne.NewMenuItem("Save", func() { sf() }),
				fyne.NewMenuItem("Export", func() { xf() }),
				fyne.NewMenuItem("Change Theme", func() { cf() }),
				fyne.NewMenuItem("Quit", func() { qf() }),
			),
			fyne.NewMenu(
				"Edit",
				fyne.NewMenuItem("Cut", func() {
					edit.TypedShortcut(
						&fyne.ShortcutCut{
							Clipboard: w.Clipboard(),
						},
					)
					info.SetText("Cut text.")
				}),
				fyne.NewMenuItem("Copy", func() {
					edit.TypedShortcut(
						&fyne.ShortcutCopy{
							Clipboard: w.Clipboard(),
						},
					)
					info.SetText("Copy text.")
				}),
				fyne.NewMenuItem("Paste", func() {
					edit.TypedShortcut(
						&fyne.ShortcutPaste{
							Clipboard: w.Clipboard(),
						},
					)
					info.SetText("Paste text.")
				}),
			),
		)
	}
	createToolbar := func() *widget.Toolbar {
		return widget.NewToolbar(
			widget.NewToolbarAction(
				theme.DocumentCreateIcon(),
				func() { nf() },
			),
			widget.NewToolbarAction(
				theme.NavigateNextIcon(),
				func() { wf() },
			),
			widget.NewToolbarAction(
				theme.SearchIcon(),
				func() { ff() },
			),
			widget.NewToolbarAction(
				theme.DocumentSaveIcon(),
				func() { sf() },
			),
		)
	}

	mb := createMenubar()
	tb := createToolbar()

	fc := widget.NewVBox(
		tb,
		widget.NewForm(
			widget.NewFormItem("FIND", find),
		),
		widget.NewHBox(cbtn, wbtn, fbtn, ibtn, sbtn, xbtn),
	)

	w.SetMainMenu(mb)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(fc, info, nil, nil),
			fc,
			info,
			sc,
		),
	)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
