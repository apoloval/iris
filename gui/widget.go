package gui

import "github.com/apoloval/karen/gfx"

// Widget is an abstraction of a GUI component
type Widget interface {
	Draw(canvas gfx.Canvas)
}

// WidgetProps are the basic properties of a widget
type WidgetProps struct {
	Align   gfx.Align
	Padding gfx.Padding
}

// Apply the given widget properties
func (p *WidgetProps) Apply(props []WidgetProp) {
	for _, prop := range props {
		prop(p)
	}
}

// WidgetProp is a set policy of widget properties
type WidgetProp func(p *WidgetProps)

// WithAlign is a setter of widget alignment
func WithAlign(align gfx.Align) WidgetProp {
	return func(p *WidgetProps) { p.Align = align }
}

// WithPadding is a setter of widget padding
func WithPadding(padding gfx.Padding) WidgetProp {
	return func(p *WidgetProps) { p.Padding = padding }
}

// TextProps are the properties of a text-based widget
type TextProps struct {
	FontType  gfx.TextFontType
	FontSize  gfx.TextFontSize
	FontColor gfx.Color
}

// Apply the given text properties
func (p *TextProps) Apply(props []TextProp) {
	for _, prop := range props {
		prop(p)
	}
}

// TextProp is a set policy of text properties
type TextProp func(*TextProps)

// WithFontType defines the text parameters of a label widget
func WithFontType(ft gfx.TextFontType) TextProp {
	return func(tp *TextProps) {
		tp.FontType = ft
	}
}

// WithFontSize defines the text parameters of a label widget
func WithFontSize(fs gfx.TextFontSize) TextProp {
	return func(tp *TextProps) {
		tp.FontSize = fs
	}
}

// WithFontColor defines the text parameters of a label widget
func WithFontColor(col gfx.Color) TextProp {
	return func(tp *TextProps) {
		tp.FontColor = col
	}
}
