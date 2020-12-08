package sdl

import (
	"fmt"
	"runtime"

	"github.com/apoloval/karen/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Engine is the Engine graphics engine
type Engine struct {
	window *sdl.Window
	screen *sdl.Surface
	fonts  map[string]*ttf.Font
}

// NewEngine instantiates the SDL graphics engine
func NewEngine(cfg *gfx.Config) (*Engine, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if err := ttf.Init(); err != nil {
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

	surface, err := window.GetSurface()
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		window: window,
		screen: surface,
		fonts:  make(map[string]*ttf.Font),
	}
	return engine, nil
}

// Canvas returns an SDL canvas
func (e *Engine) Canvas() gfx.Canvas {
	return &Canvas{
		engine: e,
	}
}

// RenderText renders the given text
func (e *Engine) RenderText(text string, params gfx.RenderTextParams) (gfx.RenderedText, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	font, err := e.font(params.Font, params.Size)
	if err != nil {
		return nil, err
	}

	surface, err := font.RenderUTF8Blended(text, sdl.Color(params.Color))
	if err != nil {
		return nil, err
	}

	return &renderedText{surface}, nil
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

func (e *Engine) font(fontType gfx.TextFontType, fontSize gfx.TextFontSize) (*ttf.Font, error) {
	key := fmt.Sprintf("%s:%d", fontType, fontSize)

	if f, ok := e.fonts[key]; ok {
		return f, nil
	}

	f, err := ttf.OpenFont(e.fontPath(fontType), int(fontSize))
	if err != nil {
		return nil, err
	}

	e.fonts[key] = f
	return f, nil
}

func (e *Engine) fontPath(fontType gfx.TextFontType) string {
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("C:\\WINDOWS\\FONTS\\%s.TTF", fontType)
	default:
		panic(fmt.Errorf("unknown platform: %s", runtime.GOOS))
	}
}

type renderedText struct {
	surface *sdl.Surface
}
