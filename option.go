package karen

import (
	"image"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/app"
)

// AppOption is an application option
type AppOption func(*AppConfig) error

// WidgetOption is a widget option
type WidgetOption func(*app.DrawProps)

//FontSize is a widget option to set the size of the text font
func FontSize(size gfx.TextFontSize) WidgetOption {
	return func(p *app.DrawProps) {
		p.DefineFontSize(size)
	}
}

// FontColor is a widget option to set the color of the text font
func FontColor(col gfx.Color) WidgetOption {
	return func(p *app.DrawProps) {
		p.DefineFontColor(col)
	}
}

// Expand is a widget option to expand the widget size
func Expand(size image.Point) WidgetOption {
	return func(p *app.DrawProps) {
		p.DefineExpand(size)
	}
}

// Align is a widget option to set the alignment
func Align(a gfx.Align) WidgetOption {
	return func(p *app.DrawProps) {
		p.DefineAlign(a)
	}
}
