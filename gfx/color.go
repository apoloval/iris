package gfx

import "image/color"

// Color represents a color
type Color color.RGBA

// Color constant definitions
var (
	ColorBlack    = Color{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	ColorBlue     = Color{R: 0x00, G: 0x00, B: 0x11, A: 0xff}
	ColorGreen    = Color{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
	ColorGrey     = Color{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	ColorGreyDark = Color{R: 0x40, G: 0x40, B: 0x40, A: 0xff}
	ColorRed      = Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	ColorWhite    = Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)
