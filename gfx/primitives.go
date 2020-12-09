package gfx

// Pos is the position of an object in the screen
type Pos struct {
	X int
	Y int
}

// Add this position to other
func (p Pos) Add(other Pos) Pos {
	return Pos{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Sub subtracts this position to other
func (p Pos) Sub(other Pos) Pos {
	return Pos{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// Size is the size of an object in the screen
type Size struct {
	W int
	H int
}

// ToRect converts this size into a rectangle at position zero.
func (s Size) ToRect() Rect {
	return Rect{Size: s}
}

// Rect is a rectangle of the screen
type Rect struct {
	Pos  Pos
	Size Size
}

// TopLeft is the position of the top-left corner of the rectangle
func (r Rect) TopLeft() Pos {
	return Pos{X: r.Pos.X, Y: r.Pos.Y}
}

// TopRight is the position of the top-right corner of the rectangle
func (r Rect) TopRight() Pos {
	return Pos{X: r.Pos.X + r.Size.W, Y: r.Pos.Y}
}

// BottomLeft is the position of the bottom-left corner of the rectangle
func (r Rect) BottomLeft() Pos {
	return Pos{X: r.Pos.X, Y: r.Pos.Y + r.Size.H}
}

// BottomRight is the position of the bottom-right corner of the rectangle
func (r Rect) BottomRight() Pos {
	return Pos{X: r.Pos.X + r.Size.W, Y: r.Pos.Y + r.Size.H}
}
