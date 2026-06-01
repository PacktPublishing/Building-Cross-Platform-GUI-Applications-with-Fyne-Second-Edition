package main

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func isImage(file fyne.URI) bool {
	ext := strings.ToLower(file.Extension())

	return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif"
}

func filterImages(files []fyne.URI) []fyne.URI {
	images := []fyne.URI{}

	for _, file := range files {
		if isImage(file) {
			images = append(images, file)
		}
	}

	return images
}

func makeImageGrid(images []fyne.URI) fyne.CanvasObject {
	items := []fyne.CanvasObject{}

	for _, item := range images {
		img := makeImageItem(item)
		items = append(items, img)
	}

	return widget.NewGridWrap(func() int {
		return len(items)
	}, func() fyne.CanvasObject {
		return makeImageItem(nil)
	}, func(id widget.GridWrapItemID, o fyne.CanvasObject) {
		c := o.(*fyne.Container)
		img := c.Objects[0].(*canvas.Image)
		text := c.Objects[3].(*canvas.Text)

		u := images[id]
		text.Text = u.Name()
		text.Refresh()

		loaded := loadImage(u).(*canvas.Image)
		img.Image = loaded.Image
		img.Refresh()
	})
}

func makeStatus(dir fyne.ListableURI, images []fyne.URI) fyne.CanvasObject {
	status := fmt.Sprintf("Directory %s, %d items", dir.Name(), len(images))
	return canvas.NewText(status, color.Gray{128})
}

func makeUI(dir fyne.ListableURI) fyne.CanvasObject {
	list, err := dir.List()
	if err != nil {
		log.Println("Error listing directory", err)
	}
	images := filterImages(list)
	status := makeStatus(dir, images)
	content := container.NewScroll(makeImageGrid(images))
	return container.NewBorder(nil, status, nil, nil, status, content)
}
