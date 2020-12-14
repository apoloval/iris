package app

import (
	"image"

	"github.com/apoloval/karen/gfx"
)

// Layout is a policy that determines how widgets are emplaced in the screen
type Layout interface {
	// Available returns the available space for the next widget
	Available(req image.Point) image.Rectangle

	// Used returns the space used by the whole layout
	Used() image.Rectangle

	// Next indicates the current widget was emplaced and how much space it required
	Next(size image.Point)
}

// LayoutProps are the properties of a layout
type LayoutProps struct {
	Padding int
	Expand  int
}

// VerticalLayout is a layout policy that emplaces the widgets vertically
func VerticalLayout(cursor image.Rectangle, props LayoutProps) Layout {
	used := image.Rectangle{
		Min: cursor.Min,
		Max: cursor.Min,
	}
	cursor.Min = cursor.Min.Add(image.Pt(0, props.Padding))
	cursor.Max = cursor.Max.Sub(image.Pt(0, props.Padding))
	return &verticalLayout{
		cursor: cursor,
		used:   used,
		props:  props,
	}
}

type verticalLayout struct {
	cursor image.Rectangle
	used   image.Rectangle
	props  LayoutProps
}

func (l *verticalLayout) Available(req image.Point) image.Rectangle {
	if req == image.ZP {
		return l.cursor
	}
	if l.props.Expand > req.X {
		req.X = l.props.Expand
	}
	return gfx.RectReduce(l.cursor, req)
}

func (l *verticalLayout) Used() image.Rectangle {
	return l.used
}

func (l *verticalLayout) Next(size image.Point) {
	if size.X > l.used.Dx() {
		l.used.Max.X = l.used.Min.X + size.X
	}
	l.used.Max.Y += size.Y + 2*l.props.Padding
	l.cursor.Min.Y += size.Y + 2*l.props.Padding
}

// HorizontalLayout is a layout policy that emplaces the widgets horizontally
func HorizontalLayout(cursor image.Rectangle, props LayoutProps) Layout {
	used := image.Rectangle{
		Min: cursor.Min,
		Max: cursor.Min,
	}
	cursor.Min = cursor.Min.Add(image.Pt(props.Padding, 0))
	cursor.Max = cursor.Max.Sub(image.Pt(props.Padding, 0))
	return &horizontalLayout{
		cursor: cursor,
		used:   used,
		props:  props,
	}
}

type horizontalLayout struct {
	cursor image.Rectangle
	used   image.Rectangle
	props  LayoutProps
}

func (l *horizontalLayout) Available(req image.Point) image.Rectangle {
	if req == image.ZP {
		return l.cursor
	}
	if l.props.Expand > req.Y {
		req.Y = l.cursor.Dy()
	}
	return gfx.RectReduce(l.cursor, req)
}

func (l *horizontalLayout) Used() image.Rectangle {
	return l.used
}

func (l *horizontalLayout) Next(size image.Point) {
	if size.Y > l.used.Dy() {
		l.used.Max.Y = l.used.Min.Y + size.Y
	}
	l.used.Max.X += size.X + 2*l.props.Padding
	l.cursor.Min.X += size.X + 2*l.props.Padding
}

// LayoutStack is a stack of layouts
type LayoutStack struct {
	stack []Layout
}

// NewLayoutStack initializes a new cursor stack
func NewLayoutStack() LayoutStack {
	return LayoutStack{
		stack: make([]Layout, 0, 256),
	}
}

// Push a new layout into the stack
func (s *LayoutStack) Push(l Layout) {
	s.stack = append(s.stack, l)
}

// Pop a layout from the stack
func (s *LayoutStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

// Top returns the top layout of the stack
func (s *LayoutStack) Top() Layout {
	if len(s.stack) == 0 {
		panic("obtaining top from an empty layout stack")
	}
	return s.stack[len(s.stack)-1]
}

// Reset cleans up the stack the initializes it with the given layout as top
func (s *LayoutStack) Reset(l Layout) {
	s.stack[0] = l
	s.stack = s.stack[:1]
}
