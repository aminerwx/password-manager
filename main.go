package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createUI() (*widget.Label, *widget.Entry) {
	out, in := widget.NewLabel("Hello World!"), widget.NewEntry()
	in.OnChanged = func(s string) {
		out.SetText("Hello " + s + "!")
	}
	return out, in
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")
	w.Resize(fyne.NewSize(680, 420))
	w.SetMaster()
	w.SetContent(container.NewVBox(createUI()))
	w.Show()

	a.Run()
}
