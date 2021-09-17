package request

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Header struct {
	Id       int32
	Key      string
	Value    string
	Selected bool
}

type ReqTab struct {
	RequestType string
	URL         string
	Headers     []Header
}

func BuildWindow() fyne.CanvasObject {
	reqTabs := container.NewDocTabs(buildTab())
	reqTabs.CreateTab = func() *container.TabItem {
		return buildTab()
	}
	return container.NewMax(reqTabs)
}

func buildTab() *container.TabItem {
	tab := &ReqTab{}
	methodSelector := widget.NewSelect([]string{"GET", "POST"}, func(value string) {
		tab.RequestType = value
	})
	methodSelector.SetSelected("GET")

	urlBind := binding.BindString(&tab.URL)
	urlEntry := widget.NewEntryWithData(urlBind)
	urlEntry.SetPlaceHolder("Enter Request URL")

	saveBtn := widget.NewButton("Send", func() {})
	saveBtn.Importance = widget.HighImportance
	urlBox := container.NewBorder(nil, nil, methodSelector, saveBtn, urlEntry)

	c1 := widget.NewCheck("", func(value bool) {})
	c2 := widget.NewCheck("", func(value bool) {})
	d1 := widget.NewToolbar(
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
		}),
	)
	d2 := widget.NewToolbar(
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
		}),
	)

	e1 := widget.NewEntry()
	e1.SetPlaceHolder("name")
	e2 := widget.NewEntry()
	e3 := widget.NewEntry()
	e4 := widget.NewEntry()

	r := container.NewBorder(nil, nil, c1, d1, container.NewGridWithColumns(2, e1, e3))
	s := container.NewBorder(nil, nil, c2, d2, container.NewGridWithColumns(2, e2, e4))
	params := container.NewVBox(r, s)

	requestDataTabs := container.NewAppTabs(
		container.NewTabItem("Params", params),
		container.NewTabItem("Headers", widget.NewLabel("World!")),
		container.NewTabItem("Body", widget.NewLabel("World!")),
	)

	requestArea := container.NewVBox(urlBox, requestDataTabs)
	text := canvas.NewText("Click Send to make request", color.Gray{128})

	split := container.NewVSplit(requestArea, text)
	return container.NewTabItem("New Request", split)
}
