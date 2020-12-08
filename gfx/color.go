package gfx

import "image/color"

// Color represents a color
type Color color.RGBA

// Color constant definitions
var (
	ColorWhite = Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)
