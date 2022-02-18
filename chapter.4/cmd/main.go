package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	r := widget.NewRadio(
		[]string{"One", "Two", "Three"},
		func(s string) {
			if s == "" {
				l.SetText("not selected.")
			} else {
				l.SetText("selected: " + s)
			}
		},
	)
	w.SetContent(
		widget.NewVBox(
			l,
			r,
		),
	)

	w.ShowAndRun()
}
