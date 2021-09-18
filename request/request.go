package request

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

type Header struct {
	Id       string
	Key      string
	Value    string
	Selected bool
}

type Req struct {
	RequestType string
	URL         string
	Headers     []*Header
}

func BuildWindow() fyne.CanvasObject {
	reqTabs := container.NewDocTabs(buildTab())
	reqTabs.CreateTab = func() *container.TabItem {
		return buildTab()
	}
	return container.NewMax(reqTabs)
}

func buildTab() *container.TabItem {
	req := &Req{}
	methodSelector := widget.NewSelect([]string{"GET", "POST"}, func(value string) {
		req.RequestType = value
	})
	methodSelector.SetSelected("GET")

	urlBind := binding.BindString(&req.URL)
	urlEntry := widget.NewEntryWithData(urlBind)
	urlEntry.SetPlaceHolder("Enter Request URL")

	saveBtn := widget.NewButton("Send", func() {})
	saveBtn.Importance = widget.HighImportance
	urlBox := container.NewBorder(nil, nil, methodSelector, saveBtn, urlEntry)

	headers := container.NewVBox()
	addHeader(req, headers)

	requestDataTabs := container.NewAppTabs(
		container.NewTabItem("Params", widget.NewLabel("World!")),
		container.NewTabItem("Headers", headers),
		container.NewTabItem("Body", widget.NewLabel("World!")),
	)

	requestArea := container.NewVBox(urlBox, requestDataTabs)
	text := canvas.NewText("Click Send to make request", color.Gray{128})

	split := container.NewVSplit(requestArea, text)
	return container.NewTabItem("New Request", split)
}

func addHeader(req *Req, parentContainer *fyne.Container) {
	id := uuid.New().String()
	header := &Header{Id: id}

	check := widget.NewCheck("", func(value bool) {
		header.Selected = value
	})
	deleteBtn := widget.NewToolbar()

	keyBind := binding.BindString(&header.Key)
	key := widget.NewEntryWithData(keyBind)
	key.SetPlaceHolder("Key")
	key.Validator = nil

	key.OnChanged = func(s string) {
		if len(s) > 0 {
			check.SetChecked(true)
		} else {
			check.SetChecked(false)
		}
		if req.Headers[len(req.Headers)-1].Id == id {
			addHeader(req, parentContainer)
		}
	}

	valueBind := binding.BindString(&header.Value)
	value := widget.NewEntryWithData(valueBind)
	value.SetPlaceHolder("Value")
	value.Validator = nil
	headerRow := container.NewBorder(nil, nil, check, deleteBtn, container.NewGridWithColumns(2, key, value))

	deleteBtn.Append(
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			if len(req.Headers) > 1 {
				parentContainer.Remove(headerRow)
				for i, h := range req.Headers {
					if h.Id != id {
						continue
					}
					copy(req.Headers[i:], req.Headers[i+1:])
					req.Headers[len(req.Headers)-1] = nil
					req.Headers = req.Headers[:len(req.Headers)-1]
					return
				}
			}
		}),
	)
	parentContainer.Add(headerRow)
	req.Headers = append(req.Headers, header)
}
