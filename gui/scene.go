package gui

import "github.com/apoloval/karen/gfx"

// Scene is the parent of any GUI component tree
type Scene struct {
	widgets []Widget
}

// NewScene instantiates a new GUI scene
func NewScene() *Scene {
	return &Scene{}
}

// Add a widget to the scene
func (s *Scene) Add(w Widget) {
	s.widgets = append(s.widgets, w)
}

// Draw the scene in the given canvas
func (s *Scene) Draw(canvas gfx.Canvas) {
	if s == nil {
		return
	}

	for _, w := range s.widgets {
		w.Draw(canvas)
	}
	canvas.Flush()
}
