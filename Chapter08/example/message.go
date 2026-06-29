package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	myName        = "Me"
	messageIndent = 20
)

type message struct {
	widget.BaseWidget

	text, from string
}

func newMessage(text, name string) *message {
	m := &message{text: text, from: name}
	m.ExtendBaseWidget(m)
	return m
}

func (m *message) CreateRenderer() fyne.WidgetRenderer {
	text := widget.NewLabel(m.text)
	text.Wrapping = fyne.TextWrapWord

	r := &messageRender{msg: m, bg: &canvas.Rectangle{},
		txt: text}
	r.Refresh()
	return r
}

type messageRender struct {
	msg *message

	bg  *canvas.Rectangle
	txt *widget.Label
}

func (r *messageRender) messageMinSize(s fyne.Size) fyne.Size {
	fitSize := s.Subtract(fyne.NewSize(messageIndent, 0))
	r.txt.Resize(fitSize.Max(r.txt.MinSize())) // have the wrap code run
	return r.txt.MinSize()
}

func (r *messageRender) Layout(s fyne.Size) {
	itemSize := r.messageMinSize(s)
	itemSize = itemSize.Max(fyne.NewSize(s.Width-messageIndent, s.Height))

	bgPos := fyne.NewPos(0, 0)
	if r.msg.from == myName {
		bgPos = fyne.NewPos(s.Width-itemSize.Width, 0)
	}

	r.txt.Move(bgPos)
	r.bg.Resize(itemSize)
	r.bg.Move(bgPos)
}

func (r *messageRender) MinSize() fyne.Size {
	itemSize := r.messageMinSize(r.msg.Size())
	return itemSize.Add(fyne.NewSize(messageIndent, 0))
}

func (r *messageRender) Refresh() {
	if r.msg.from == myName {
		r.txt.Alignment = fyne.TextAlignTrailing
		r.bg.FillColor = &color.NRGBA{0x28, 0x9d, 0xf2, 0xef}
	} else {
		r.txt.Alignment = fyne.TextAlignLeading
		r.bg.FillColor = &color.NRGBA{0x8f, 0xc5, 0x6b, 0xef}
	}
	r.txt.Refresh()
	r.bg.Refresh()
}

func (r *messageRender) Destroy() {
}

func (r *messageRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.txt}
}
