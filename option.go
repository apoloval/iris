package iris

import (
	"github.com/apoloval/iris/gfx"
	"github.com/apoloval/iris/internal/app"
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

// Align is a widget option to set the alignment
func Align(a gfx.Align) WidgetOption {
	return func(p *app.DrawProps) {
		p.DefineAlign(a)
	}
}

// LayoutOption is a layout option
type LayoutOption func(*app.LayoutProps)

// Padding is a layout option to set the padding of the widgets
func Padding(pixels int) LayoutOption {
	return func(p *app.LayoutProps) {
		p.Padding = pixels
	}
}

// Expand is a layout option to expand the widgets to the maximum space
func Expand(pixels int) LayoutOption {
	return func(p *app.LayoutProps) {
		p.Expand = pixels
	}
}
