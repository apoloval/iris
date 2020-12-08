package gui

import "github.com/apoloval/karen/gfx"

// Widget is an abstraction of a GUI component
type Widget interface {
	Draw(canvas gfx.Canvas)
}
