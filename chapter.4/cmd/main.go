package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	s := widget.NewSlider(0.0, 100.0)
	b := widget.NewButton("Check", func() {
		l.SetText("value: " + strconv.Itoa(int(s.Value)))
	})

	w.SetContent(
		widget.NewVBox(
			l,
			s,
			b,
		),
	)

	w.ShowAndRun()
}
