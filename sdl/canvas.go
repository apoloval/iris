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
func (c *Canvas) DrawText(text gfx.RenderedText, align gfx.Align) {
	switch t := text.(type) {
	case *renderedText:
		c.drawText(t, align)
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

func (c *Canvas) drawText(text *renderedText, align gfx.Align) {
	srcRect := sdl.Rect{
		W: text.surface.W,
		H: text.surface.H,
	}
	dstRect := sdl.Rect{
		X: int32(align.CalculateX(int(srcRect.W), int(c.engine.screen.W))),
		Y: int32(align.CalculateY(int(srcRect.H), int(c.engine.screen.H))),
		W: srcRect.W,
		H: srcRect.H,
	}
	if err := text.surface.Blit(&srcRect, c.engine.screen, &dstRect); err != nil {
		panic(err)
	}
}
