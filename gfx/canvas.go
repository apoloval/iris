package gfx

// Canvas is an abstraction for the drawing actions over a screen window
type Canvas interface {
	// DrawText draws the previously rendered text
	DrawText(dest Rect, text RenderedText)

	// Engine returns the graphics engine for this canvas
	Engine() Engine

	// Flush the pending drawing operations and render the final canvas
	Flush()

	// Size returns the size of this canvas
	Size() Size
}
