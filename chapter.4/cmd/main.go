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

	b := widget.NewButton("Click", func() {
		dialog.ShowConfirm(
			"Alert",
			"Please check 'YES'!",
			func(f bool) {
				if f {
					l.SetText("OK, thank you !")
				} else {
					l.SetText("oh no...")
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
