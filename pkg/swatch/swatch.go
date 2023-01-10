package swatch

import (
	"image/color"
	"pixl/pkg/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (swatch *Swatch) SetColor(c color.Color) {
	swatch.Color = c
	swatch.Refresh()
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}

func NewSwatch(
	state *apptype.State,
	color color.Color,
	swatchIndex int,
	clickHandler func(s *Swatch),
) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	swatch.ExtendBaseWidget(swatch)

	return swatch
}
