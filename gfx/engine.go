package gfx

import (
	"image"

	"github.com/apoloval/karen/internal/io"
)

// Engine is a graphics engine
type Engine interface {
	// BeginFrame begins a new screen frame using this engine
	BeginFrame()

	// EndFrame ends the current screen frame
	EndFrame()

	// ScreenDims returns the dimensions of the screen
	ScreenDims() image.Point

	// PollEvents polls IO events from this engine and updates the corresponding IO state with them
	PollEvents(s *io.State) error

	// TextDims calculates the dimensions of the given text and its parameters
	TextDims(text string, params RenderTextParams) (image.Point, error)

	// Apply the given render actions
	Apply(actions []DrawAction) error
}
