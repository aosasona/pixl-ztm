package ui

import (
	"pixl/pkg/apptype"
	"pixl/pkg/swatch"

	"fyne.io/fyne/v2"
)

type App struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
