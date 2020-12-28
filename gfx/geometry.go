package gfx

import "image"

// AsRectSize returns the rectangle at (0, 0) interpreting a point as size
func AsRectSize(size image.Point) image.Rectangle {
	return image.Rectangle{
		Min: image.Pt(0, 0),
		Max: size,
	}
}

// RectReduce reduces `r` to `size` unless it was already smaller
func RectReduce(r image.Rectangle, size image.Point) image.Rectangle {
	return AsRectSize(size).Add(r.Min).Intersect(r)
}

// RectAddPadding adds the given padding to `r`
func RectAddPadding(r image.Rectangle, padding int) image.Rectangle {
	p := image.Pt(padding, padding)
	return image.Rectangle{
		Min: r.Min.Sub(p),
		Max: r.Max.Add(p),
	}
}

// Align describes the widget alignment
type Align int

// Alignment constants
const (
	AlignTopLeft     Align = 0
	AlignTop         Align = 2
	AlignTopRight    Align = 3
	AlignLeft        Align = 4
	AlignCenter      Align = 5
	AlignRight       Align = 6
	AlignBottomLeft  Align = 7
	AlignBottom      Align = 8
	AlignBottomRight Align = 9
)

// Apply this alignment to the given rectangles
func (a Align) Apply(src image.Point, dest image.Rectangle) image.Rectangle {
	sizeDiff := dest.Size().Sub(src)
	switch a {
	case AlignTopLeft, AlignLeft, AlignBottomLeft:
		dest.Max.X -= sizeDiff.X
	case AlignTop, AlignCenter, AlignBottom:
		dest.Min.X += sizeDiff.X / 2
		dest.Max.X -= sizeDiff.X / 2
	case AlignTopRight, AlignRight, AlignBottomRight:
		dest.Min.X += sizeDiff.X
	}

	switch a {
	case AlignTopLeft, AlignTop, AlignTopRight:
		dest.Max.Y -= sizeDiff.Y
	case AlignLeft, AlignCenter, AlignRight:
		dest.Min.Y += sizeDiff.Y / 2
		dest.Max.Y -= sizeDiff.Y / 2
	case AlignBottomLeft, AlignBottom, AlignBottomRight:
		dest.Min.Y += sizeDiff.Y
	}

	return dest
}
