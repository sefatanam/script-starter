package ui

import (
	"fmt"
	"fynegui/constant"
	"fynegui/lib"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func HomeWindow(appInstance fyne.App, scriptRefs *[]lib.App) fyne.Window {

	homeWindow := appInstance.NewWindow("Script Starter")
	homeWindow.Resize(fyne.NewSize(constant.Width, constant.Height))
	homeWindow.SetFixedSize(true)

	addedAppsLabel := widget.NewLabel("Added Scripts")
	addedAppsLabel.TextStyle = fyne.TextStyle{Bold: true}

	addAppButton := widget.NewButton("Add Script", func() {
		ShowFormWindow(homeWindow)
	})

	headerRow := container.NewHBox(addedAppsLabel, layout.NewSpacer(), addAppButton)

	addedAppsTable := widget.NewTable(
		func() (rows int, cols int) {
			return len(*scriptRefs) + 1, 4
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("fallback content"))
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Row == 0 { // Header row
				switch tci.Col {
				case 0:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("#")}
				case 1:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Name")}
				case 2:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Source")}
				case 3:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Actions")}
				}
			} else {
				record := (*scriptRefs)[tci.Row-1] // Adjust index because of header row
				switch tci.Col {
				case 0:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(record.GetDetail("Id"))}
				case 1:
					co.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(record.GetDetail("Name"))}
				case 2:

					sourceLabel := widget.NewLabelWithStyle(lib.ShortenSentenceToFixChar(record.GetDetail("Source"), 30), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
					sourceLabel.Resize(fyne.NewSize(150, 200)) // Adjust the width as needed
					co.(*fyne.Container).Objects = []fyne.CanvasObject{sourceLabel}
				case 3:
					startIcon := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
						fmt.Printf("Start action for row %d\n", tci.Row)
						record.Run()
					})
					// Comment the close icon
					/* closeIcon := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
						fmt.Printf("Close action for row %d\n", tci.Row)
					}) */
					co.(*fyne.Container).Objects = []fyne.CanvasObject{startIcon}
				}
			}
		},
	)

	const minWidth float32 = 150
	addedAppsTable.SetColumnWidth(0, 50)         // Set a fixed width for the first column
	addedAppsTable.SetColumnWidth(1, minWidth)   // Set a fixed width for the second column
	addedAppsTable.SetColumnWidth(2, minWidth*2) // Set a fixed width for the third column
	addedAppsTable.SetColumnWidth(3, minWidth/3) // Set a fixed width for the action column

	addedAppsTable.StickyColumnCount = 0

	scrollContainer := container.NewVScroll(addedAppsTable)
	scrollContainer.SetMinSize(fyne.NewSize(constant.Width, constant.Height))

	startAllButton := widget.NewButton("Start All", func() {
		for _, app := range *scriptRefs {
			app.Run()
		}
	})
	footerRow := container.NewHBox(layout.NewSpacer(), startAllButton)

	appContainer := container.NewBorder(headerRow, footerRow, nil, nil, scrollContainer)
	homeWindow.SetContent(appContainer)

	return homeWindow
}
