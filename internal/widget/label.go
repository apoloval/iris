package widget

import (
	"image"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/app"
)

// DefaultLabelFontSize is the default font size for label widgets
const DefaultLabelFontSize gfx.TextFontSize = 14

// DefaultLabelFontColor is the default font color for label widgets
var DefaultLabelFontColor gfx.Color = gfx.ColorWhite

// Label emplaces a label widget with the given text.
// Returns true if the label is mouse focused.
func Label(s *app.State, wid uint, text string) bool {
	textParams := gfx.RenderTextParams{
		Font:  gfx.TextFontTypeArial,
		Size:  s.DrawProps.FontSize(DefaultLabelFontSize),
		Color: s.DrawProps.FontColor(DefaultLabelFontColor),
	}

	size, err := s.Engine.TextDims(text, textParams)
	if err != nil {
		panic(err)
	}

	dest := image.Rect(
		s.Cursor.X,
		s.Cursor.Y,
		s.Cursor.X+size.X,
		s.Cursor.Y+size.Y,
	)
	mouseFocused := s.IO.MouseIn(dest)

	s.DrawList.Append(gfx.DrawText{
		Dest:   dest,
		Text:   text,
		Params: textParams,
	})

	s.Cursor.Y += size.Y
	s.DrawProps.Reset()

	return mouseFocused
}
