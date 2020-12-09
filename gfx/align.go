package gfx

// Align is an object position alignment policy
type Align func(src Rect, dst Rect) Rect

// AlignLeft is an alignment policy to place the source at the left-middle of the destination
func AlignLeft(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X,
			Y: dst.Pos.Y + (dst.Size.H-src.Size.H)/2,
		},
		Size: src.Size,
	}
}

// AlignRight is an alignment policy to place the source at the right-middle of the destination
func AlignRight(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + dst.Size.W - src.Size.W,
			Y: dst.Pos.Y + (dst.Size.H-src.Size.H)/2,
		},
		Size: src.Size,
	}
}

// AlignTop is an alignment policy to place the source at the top-center of the destination
func AlignTop(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + (dst.Size.W-src.Size.W)/2,
			Y: dst.Pos.Y,
		},
		Size: src.Size,
	}
}

// AlignBottom is an alignment policy to place the source at the bottom-center of the destination
func AlignBottom(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + (dst.Size.W-src.Size.W)/2,
			Y: dst.Pos.Y + dst.Size.H - src.Size.H,
		},
		Size: src.Size,
	}
}

// AlignTopLeft is an alignment policy to place the source at the top-left of the destination
func AlignTopLeft(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X,
			Y: dst.Pos.Y,
		},
		Size: src.Size,
	}
}

// AlignTopRight is an alignment policy to place the source at the top-right of the destination
func AlignTopRight(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + dst.Size.W - src.Size.W,
			Y: dst.Pos.Y,
		},
		Size: src.Size,
	}
}

// AlignBottomLeft is an alignment policy to place the source at the bottom-left of the destination
func AlignBottomLeft(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X,
			Y: dst.Pos.Y + dst.Size.H - src.Size.H,
		},
		Size: src.Size,
	}
}

// AlignBottomRight is an alignment policy to place the source at the bottom-right of the destination
func AlignBottomRight(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + dst.Size.W - src.Size.W,
			Y: dst.Pos.Y + dst.Size.H - src.Size.H,
		},
		Size: src.Size,
	}
}

// AlignCenter is an alignment policy to place the source at the center of the destination
func AlignCenter(src Rect, dst Rect) Rect {
	return Rect{
		Pos: Pos{
			X: dst.Pos.X + (dst.Size.W-src.Size.W)/2,
			Y: dst.Pos.Y + (dst.Size.H-src.Size.H)/2,
		},
		Size: src.Size,
	}
}
