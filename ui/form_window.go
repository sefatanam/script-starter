package ui

import (
	"fynegui/constant"
	"fynegui/lib"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowFormWindow(previousWindowRef fyne.Window) {

	previousWindowRef.Hide()

	newWindow := fyne.CurrentApp().NewWindow("Submit Script")
	newWindow.Resize(fyne.NewSize(constant.Width, constant.Height*1.2))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter script name")

	sourceEntry := widget.NewEntry()
	sourceEntry.SetPlaceHolder("Browse file path")
	sourceEntry.TextStyle.Bold = true
	sourceEntry.TextStyle.Monospace = true
	sourceEntry.Disable()

	sourceButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFolderOpen(func(file fyne.ListableURI, err error) {
			if file == nil {
				return
			} else {
				sourceEntry.SetText(string(file.Path()))
			}
		}, newWindow)

		fileDialog.Show()

	})

	sourceContainer := container.New(layout.NewFormLayout(), sourceButton, sourceEntry)

	form := container.NewVBox(
		container.NewVBox(widget.NewLabel("Name"), nameEntry),
		layout.NewSpacer(),
		widget.NewLabel("Browse Path"),
		sourceContainer,
		layout.NewSpacer(),
	)
	submitButton := widget.NewButton("Submit", func() {
		script := lib.Script{
			Id:     0,
			Name:   nameEntry.Text,
			Source: sourceEntry.Text,
		}
		script.AddScript()
		newWindow.Close()
		previousWindowRef.Content().Refresh()
		previousWindowRef.Show()

	})

	cancelButton := widget.NewButton("Cancel", func() {
		newWindow.Close()
		previousWindowRef.Content().Refresh()
		previousWindowRef.Show()
	})

	buttonContainer := container.NewHBox(layout.NewSpacer(), cancelButton, submitButton)

	content := container.NewBorder(form, buttonContainer, nil, nil)
	newWindow.SetContent(content)
	newWindow.Show()
}
