package io

import (
	"image"
)

// State describes the state of the IO devices
type State struct {
	MousePos image.Point
	Quit     bool
}

// MouseIn returns true if the mouse pointer is inside the given rectangle
func (s State) MouseIn(r image.Rectangle) bool {
	return s.MousePos.In(r)
}
