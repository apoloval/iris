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

	Cursor image.Point

	FrameStart        time.Time
	FrameEnd          time.Time
	LastFrameDuration time.Duration
}

// NewState initializes a new application state
func NewState(e gfx.Engine) *State {
	return &State{
		Engine: e,
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
	s.Cursor = image.Pt(0, 0)
}

// EndFrame indicates to the app the end of the current frame
func (s *State) EndFrame() {
	s.Engine.Apply(s.DrawList)
	s.Engine.EndFrame()
	s.FrameEnd = time.Now()
	s.LastFrameDuration = s.FrameEnd.Sub(s.FrameStart)
}
