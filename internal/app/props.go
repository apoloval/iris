package app

import (
	"github.com/apoloval/iris/gfx"
)

// DrawProps are the drawing properties used by the app for the next widget
type DrawProps struct {
	fontSize  Prop
	fontColor Prop
	align     Prop
}

// Reset the drawing options
func (p *DrawProps) Reset() {
	p.fontSize.Undefine()
	p.fontColor.Undefine()
	p.align.Undefine()
}

// DefineFontSize sets the size of the text font
func (p *DrawProps) DefineFontSize(size gfx.TextFontSize) {
	p.fontSize.Define(size)
}

// FontSize returns the text font size property, or the given fallback if undefined
func (p *DrawProps) FontSize(fb gfx.TextFontSize) gfx.TextFontSize {
	return p.fontSize.Get(fb).(gfx.TextFontSize)
}

// DefineFontColor sets the color of the text font
func (p *DrawProps) DefineFontColor(col gfx.Color) {
	p.fontColor.Define(col)
}

// FontColor returns the text font color property, or the given fallback if undefined
func (p *DrawProps) FontColor(fb gfx.Color) gfx.Color {
	return p.fontColor.Get(fb).(gfx.Color)
}

// DefineAlign sets the widget alignment
func (p *DrawProps) DefineAlign(a gfx.Align) {
	p.align.Define(a)
}

// Align returns the widget alignment
func (p *DrawProps) Align(fb gfx.Align) gfx.Align {
	return p.align.Get(fb).(gfx.Align)
}

// Prop is a generic application property
type Prop struct {
	value   interface{}
	defined bool
}

// Define the property to the given value
func (p *Prop) Define(value interface{}) {
	p.value = value
	p.defined = true
}

// Undefine the property
func (p *Prop) Undefine() {
	p.defined = false
}

// Get the property value, or the given fallback if undefined
func (p *Prop) Get(fb interface{}) interface{} {
	if !p.defined {
		return fb
	}
	return p.value
}
