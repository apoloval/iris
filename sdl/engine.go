package sdl

import (
	"github.com/apoloval/karen/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Engine is the Engine graphics engine
type Engine struct {
	win *sdl.Window
}

// NewEngine instantiates the SDL graphics engine
func NewEngine(cfg *gfx.Config) (*Engine, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(
		cfg.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(cfg.ScreenWidth),
		int32(cfg.ScreenHeight),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		win: window,
	}
	return engine, nil
}

// Canvas returns an SDL canvas
func (e *Engine) Canvas() gfx.Canvas {
	return &Canvas{}
}

// WaitEvent waits until a SDL event is produced
func (e *Engine) WaitEvent() (gfx.Event, error) {
	for {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			return gfx.EventQuit{}, nil
		}
	}
}
