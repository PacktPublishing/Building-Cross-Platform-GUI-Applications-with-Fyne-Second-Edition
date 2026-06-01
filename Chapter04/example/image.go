package main

import (
	"image"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"github.com/nfnt/resize"
)

func scaleImage(img image.Image) image.Image {
	return resize.Thumbnail(320, 240, img, resize.Lanczos3)
}

func loadImage(u fyne.URI) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScaleFastest

	if u == nil {
		return img
	}

	r, err := storage.Reader(u)
	if err != nil {
		log.Println("Error reading image", err)
		return img
	}

	src, _, err := image.Decode(r)
	if err != nil {
		log.Println("Error decoding image", err)
		return img
	}

	thumb := scaleImage(src)
	img.Image = thumb
	return img
}

type itemLayout struct {
	bg, text, gradient fyne.CanvasObject
}

func (i *itemLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(160, 120)
}

func (i *itemLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	textHeight := float32(22)
	for _, o := range objs {
		if o == i.text {
			o.Move(fyne.NewPos(0, size.Height-textHeight))
			o.Resize(fyne.NewSize(size.Width, textHeight))
		} else if o == i.bg {
			o.Move(fyne.NewPos(0, size.Height-textHeight))
			o.Resize(fyne.NewSize(size.Width, textHeight))
		} else if o == i.gradient {
			o.Move(fyne.NewPos(0, size.Height-(textHeight*1.5)))
			o.Resize(fyne.NewSize(size.Width, textHeight/2))
		} else {
			o.Move(fyne.NewPos(0, 0))
			o.Resize(size)
		}
	}
}

func makeImageItem(u fyne.URI) fyne.CanvasObject {
	text := ""
	if u != nil {
		text = u.Name()
	}
	label := canvas.NewText(text, color.Gray{128})
	label.Alignment = fyne.TextAlignCenter

	bgColor := &color.NRGBA{R: 255, G: 255, B: 255, A: 224}
	bg := canvas.NewRectangle(bgColor)
	fade := canvas.NewLinearGradient(color.Transparent, bgColor, 0)
	return container.New(&itemLayout{text: label, bg: bg, gradient: fade},
		loadImage(u), bg, fade, label)
}
