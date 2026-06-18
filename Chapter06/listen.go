package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	_ = app.New()

	val := binding.NewString()
	callback := binding.NewDataListener(func() {
		str, _ := val.Get()
		fmt.Println("String changed to:", str)
	})

	val.AddListener(callback)
	_ = val.Set("new data")
}
