package sdl

import (
	"github.com/apoloval/karen/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// ToSDLRect converts a `gfx.Rect` into a `sdl.Rect`
func ToSDLRect(r gfx.Rect) sdl.Rect {
	return sdl.Rect{
		X: int32(r.Pos.X),
		Y: int32(r.Pos.Y),
		W: int32(r.Size.W),
		H: int32(r.Size.H),
	}
}

// FromSDLRect converts a `sdl.Rect` into a `gfx.Rect`
func FromSDLRect(r sdl.Rect) gfx.Rect {
	return gfx.Rect{
		Pos: gfx.Pos{
			X: int(r.X),
			Y: int(r.Y),
		},
		Size: gfx.Size{
			W: int(r.W),
			H: int(r.H),
		},
	}
}
