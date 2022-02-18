package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	sl := widget.NewSelect(
		[]string{"Eins", "Twei", "Drei"},
		func(s string) {
			l.SetText("selected: " + s)
		},
	)
	w.SetContent(
		widget.NewVBox(
			l,
			sl,
		),
	)

	w.ShowAndRun()
}
