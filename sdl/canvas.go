package sdl

import (
	"github.com/apoloval/karen/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Canvas is a SDL-based canvas implementation
type Canvas struct {
	engine *Engine
}

// DrawText draws the given text into this canvas
func (c *Canvas) DrawText(dest gfx.Rect, text gfx.RenderedText) {
	switch t := text.(type) {
	case *renderedText:
		c.drawText(dest, t)
	}
}

// Engine returns the SDL engine for this canvas
func (c *Canvas) Engine() gfx.Engine {
	return c.engine
}

// Flush the pending drawing operations and render the final canvas
func (c *Canvas) Flush() {
	if err := c.engine.window.UpdateSurface(); err != nil {
		panic(err)
	}
}

// Size is the size of this canvas
func (c *Canvas) Size() gfx.Size {
	return gfx.Size{W: int(c.engine.screen.W), H: int(c.engine.screen.H)}
}

func (c *Canvas) drawText(dest gfx.Rect, text *renderedText) {
	srcRect := sdl.Rect{
		W: text.surface.W,
		H: text.surface.H,
	}
	dstRect := ToSDLRect(dest)
	if err := text.surface.Blit(&srcRect, c.engine.screen, &dstRect); err != nil {
		panic(err)
	}
}
