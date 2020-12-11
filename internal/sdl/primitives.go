package sdl

import (
	"image"

	"github.com/veandco/go-sdl2/sdl"
)

// ToSDLRect converts a `gfx.Rect` into a `sdl.Rect`
func ToSDLRect(r image.Rectangle) sdl.Rect {
	return sdl.Rect{
		X: int32(r.Min.X),
		Y: int32(r.Min.Y),
		W: int32(r.Dx()),
		H: int32(r.Dy()),
	}
}
