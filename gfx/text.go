package gfx

import "fmt"

// TextFontType describes a type of text font
type TextFontType string

// Some standard font types
const (
	TextFontTypeArial TextFontType = "ARIAL"
)

// TextFontSize describes the size of a font in typography points (PTs)
type TextFontSize int

// RenderTextParams are the paratemers to draw text in a canvas
type RenderTextParams struct {
	Font  TextFontType
	Size  TextFontSize
	Color Color
}

// Validate the render text parameters
func (p RenderTextParams) Validate() error {
	if p.Font == "" {
		return fmt.Errorf("invalid render text params: missing font type")
	}
	if p.Size <= 0 {
		return fmt.Errorf("invalid render text params: invalid font size")
	}
	return nil
}

// RenderedText is a fragment ot text rendered by the canvas
type RenderedText interface{}
