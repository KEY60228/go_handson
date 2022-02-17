package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	c := widget.NewCheck("Check!", func(f bool) {
		if f {
			l.SetText("CHECKED!")
		} else {
			l.SetText("not checked.")
		}
	})
	w.SetContent(
		widget.NewVBox(
			l,
			c,
		),
	)

	w.ShowAndRun()
}

func total(n int) int {
	var t int
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}
