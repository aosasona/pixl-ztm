package ui

import (
	"pixl/pkg/apptype"
	"pixl/pkg/pxcanvas"
	"pixl/pkg/swatch"

	"fyne.io/fyne/v2"
)

type App struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
