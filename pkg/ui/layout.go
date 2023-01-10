package ui

import "fyne.io/fyne/v2/container"

func Setup(app *App) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
