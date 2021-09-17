package main

import (
	"github.com/sajan45/res/request"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	resApp := app.New()
	mainWindow := resApp.NewWindow("RES")
	tabContainer := request.BuildWindow()
	mainWindow.SetContent(tabContainer)
	mainWindow.Resize(fyne.NewSize(900, 640))
	mainWindow.ShowAndRun()
}
