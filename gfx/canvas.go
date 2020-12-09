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

// WithPadding applies some padding to the given canvas
func WithPadding(parent Canvas, padding Padding) Canvas {
	cs := parent.Size()
	view := Rect{
		Pos:  Pos{X: int(padding), Y: int(padding)},
		Size: Size{W: cs.W - 2*int(padding), H: cs.H - 2*int(padding)},
	}
	return &canvasView{parent, view}
}

type canvasView struct {
	parent Canvas
	view   Rect
}

func (c canvasView) DrawText(dest Rect, text RenderedText) {
	c.parent.DrawText(c.parentRelative(dest), text)
}

func (c canvasView) Engine() Engine {
	return c.parent.Engine()
}

func (c canvasView) Flush() {
	c.parent.Flush()
}

func (c canvasView) Size() Size {
	return c.view.Size
}

func (c canvasView) parentRelative(r Rect) Rect {
	return Rect{
		Pos: Pos{
			X: c.view.Pos.X + r.Pos.X,
			Y: c.view.Pos.Y + r.Pos.Y,
		},
		Size: r.Size,
	}
}
