package gfx

import "image"

// DrawAction is a rendering action passed to the renderer
type DrawAction interface {
	IsDrawAction()
}

// DrawTexture is an action to draw a image texture on the screen
type DrawTexture struct {
	Dest    image.Rectangle
	Texture Texture
}

// IsDrawAction is a function from marking interface `RenderAction`
func (DrawTexture) IsDrawAction() {}

// DrawText is an action to draw a text on the screen
type DrawText struct {
	Dest   image.Rectangle
	Text   string
	Params RenderTextParams
}

// IsDrawAction is a function from marking interface `RenderAction`
func (DrawText) IsDrawAction() {}

// DrawList is a slice of draw actions
type DrawList []DrawAction

// Append a draw action to this list
func (b *DrawList) Append(a DrawAction) {
	*b = append(*b, a)
}

// Clean this draw list
func (b *DrawList) Clean() {
	*b = (*b)[:0]
}

// Texture is an image already processed by the engine and ready to be draw in the screen
type Texture interface {
	// Size is the size of the texture
	Size() image.Point
}
