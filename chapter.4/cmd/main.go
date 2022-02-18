package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")

	e := widget.NewEntry()
	b := widget.NewButton("Click", func() {
		dialog.ShowCustomConfirm(
			"Enter message.",
			"OK",
			"Cancel",
			e,
			func(f bool) {
				if f {
					l.SetText("typed: '" + e.Text + "'.")
				} else {
					l.SetText("no message.")
				}
			},
			w,
		)
	})

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(nil, b, nil, nil), l, b,
		),
	)

	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(350, 250))
	w.ShowAndRun()
}
