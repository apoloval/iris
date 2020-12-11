package app

import "github.com/apoloval/karen/gfx"

// DrawProps are the drawing properties used by the app for the next widget
type DrawProps struct {
	fontSize  gfx.TextFontSize
	fontColor gfx.Color

	fontSizeDefined  bool
	fontColorDefined bool
}

// Reset the drawing options
func (p *DrawProps) Reset() {
	p.fontSizeDefined = false
	p.fontColorDefined = false
}

// DefineFontSize sets the size of the text font
func (p *DrawProps) DefineFontSize(size gfx.TextFontSize) {
	p.fontSize = size
	p.fontSizeDefined = true
}

// FontSize returns the text font size property, or the given fallback if undefined
func (p *DrawProps) FontSize(fb gfx.TextFontSize) gfx.TextFontSize {
	if !p.fontSizeDefined {
		return fb
	}
	return p.fontSize
}

// DefineFontColor sets the color of the text font
func (p *DrawProps) DefineFontColor(col gfx.Color) {
	p.fontColor = col
	p.fontColorDefined = true
}

// FontColor returns the text font color property, or the given fallback if undefined
func (p *DrawProps) FontColor(fb gfx.Color) gfx.Color {
	if !p.fontColorDefined {
		return fb
	}
	return p.fontColor
}
