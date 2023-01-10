package pixl

import (
	"image/color"
	"pixl/pkg/apptype"
	"pixl/pkg/pxcanvas"
	"pixl/pkg/swatch"
	"pixl/pkg/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	pixlConfig := ui.App{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&pixlConfig)

	pixlConfig.PixlWindow.ShowAndRun()
}
