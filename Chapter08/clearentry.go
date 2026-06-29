package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type clearEntry struct {
	widget.Entry
}

func newClearEntry() *clearEntry {
	e := &clearEntry{}
	e.ExtendBaseWidget(e)
	return e
}

func (s *clearEntry) TypedKey(k *fyne.KeyEvent) {
	if k.Name == fyne.KeyEscape {
		s.SetText("")
		return
	}

	s.Entry.TypedKey(k)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clear Entry")

	w.SetContent(newClearEntry())
	w.ShowAndRun()
}
