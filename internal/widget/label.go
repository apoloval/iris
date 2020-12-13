package widget

import (
	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/app"
)

// Default label constants
const (
	DefaultLabelFontSize gfx.TextFontSize = 14
	DefaultLabelAlign    gfx.Align        = gfx.AlignTopLeft
)

// DefaultLabelFontColor is the default font color for label widgets
var DefaultLabelFontColor gfx.Color = gfx.ColorWhite

// Label emplaces a label widget with the given text.
// Returns true if the label is mouse focused.
func Label(s *app.State, wid uint, text string) bool {
	textParams := gfx.RenderTextParams{
		Font:  gfx.TextFontTypeArial,
		Size:  s.DrawProps.FontSize(DefaultLabelFontSize),
		Color: s.DrawProps.FontColor(DefaultLabelFontColor),
		Align: s.DrawProps.Align(DefaultLabelAlign),
	}

	size, err := s.Engine.TextDims(text, textParams)
	if err != nil {
		panic(err)
	}

	exp := s.DrawProps.Expand(size)
	if exp.X > size.X {
		size.X = exp.X
	}
	if exp.Y > size.Y {
		size.Y = exp.Y
	}

	dest := s.Available(size)
	mouseFocused := s.IO.MouseIn(dest)

	s.DrawList.Append(gfx.DrawText{
		Dest:   dest,
		Text:   text,
		Params: textParams,
	})

	s.Next(dest.Size())
	return mouseFocused
}
