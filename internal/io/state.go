package io

import (
	"image"
)

// State describes the state of the IO devices
type State struct {
	MousePos image.Point
	Quit     bool
}

// NewState instantiates a new IO state
func NewState() State {
	return State{
		MousePos: image.Pt(-1, -1),
	}
}

// MouseIn returns true if the mouse pointer is inside the given rectangle
func (s State) MouseIn(r image.Rectangle) bool {
	return s.MousePos.In(r)
}
