package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")
	tb := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.HomeIcon(),
			func() {
				l.SetText("Select Home Icon!")
			},
		),
		widget.NewToolbarAction(
			theme.InfoIcon(),
			func() {
				l.SetText("Select Information Icon!")
			},
		),
	)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(nil, tb, nil, nil),
			l,
			tb,
		),
	)

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
