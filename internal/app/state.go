package app

import (
	"image"
	"time"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/io"
)

// State is the internal state of the application
type State struct {
	Engine    gfx.Engine
	IO        io.State
	DrawList  gfx.DrawList
	DrawProps DrawProps

	Layout LayoutStack

	FrameStart        time.Time
	FrameEnd          time.Time
	LastFrameDuration time.Duration
}

// NewState initializes a new application state
func NewState(e gfx.Engine) *State {
	return &State{
		Engine: e,
		IO:     io.NewState(),
		Layout: NewLayoutStack(),
	}
}

// BeginFrame indicates to the app the beginning of a new frame
func (s *State) BeginFrame() {
	s.FrameStart = time.Now()
	if err := s.Engine.PollEvents(&s.IO); err != nil {
		panic(err)
	}
	s.Engine.BeginFrame()
	s.DrawList.Clean()

	screenRect := image.Rectangle{
		Min: image.Pt(0, 0),
		Max: s.Engine.ScreenDims(),
	}
	layout := VerticalLayout(screenRect)
	s.Layout.Push(layout)
}

// EndFrame indicates to the app the end of the current frame
func (s *State) EndFrame() {
	s.Engine.Apply(s.DrawList)
	s.Engine.EndFrame()
	s.FrameEnd = time.Now()
	s.LastFrameDuration = s.FrameEnd.Sub(s.FrameStart)
}

// BeginLayoutH begins a horizontal layout
func (s *State) BeginLayoutH() {
	l := HorizontalLayout(s.Layout.Top().Available(image.ZP))
	s.Layout.Push(l)
}

// EndLayout ends the current layout
func (s *State) EndLayout() {
	used := s.Layout.Top().Used()
	s.Layout.Pop()
	s.Layout.Top().Next(used.Size())
}

// Available returns the available space for the current widget
func (s *State) Available(req image.Point) image.Rectangle {
	return s.Layout.Top().Available(req)
}

// Next signals the state must be prepared for the next widget in the layout
func (s *State) Next(size image.Point) {
	s.Layout.Top().Next(size)
	s.DrawProps.Reset()
}
