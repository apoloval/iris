package karen

import (
	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/app"
)

// WidgetOpt is a widget option
type WidgetOpt func(*app.DrawProps)

func FontSize(size gfx.TextFontSize) WidgetOpt {
	return func(p *app.DrawProps) {
		p.DefineFontSize(size)
	}
}

func FontColor(col gfx.Color) WidgetOpt {
	return func(p *app.DrawProps) {
		p.DefineFontColor(col)
	}
}
