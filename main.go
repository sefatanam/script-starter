package main

import (
	"fynegui/lib"
	"fynegui/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	appInstance := app.New()
	scripts := lib.GetScripts()
	ui.HomeWindow(appInstance, scripts).Show()
	appInstance.Run()
}
