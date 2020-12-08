package karen

import (
	"errors"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/gui"
	"github.com/apoloval/karen/sdl"
)

// ErrUnknownEngine is an error returned when an unknown engine is specified
var ErrUnknownEngine = errors.New("unknown GFX engine")

// App is the Karen application object.
type App struct {
	config   *Config
	graphics gfx.Engine
	scene    *gui.Scene
}

// NewApp instantiates a new application
func NewApp(opts ...Option) (*App, error) {
	cfg := defaultConfig()
	if err := applyOptions(cfg, opts); err != nil {
		return nil, err
	}

	graphics, err := newGraphics(cfg)
	if err != nil {
		return nil, err
	}

	app := &App{
		config:   cfg,
		graphics: graphics,
	}
	return app, nil
}

// NewScene instantiates a new UI scene
func (a *App) NewScene() *gui.Scene {
	s := gui.NewScene()
	a.scene = s
	return s
}

// Run this application until closed or fails
func (a *App) Run() error {
	a.scene.Draw(a.graphics.Canvas())
	for {
		ev, err := a.graphics.WaitEvent()
		if err != nil {
			return err
		}

		switch ev.(type) {
		case gfx.EventQuit:
			return nil
		}
	}
}

func newGraphics(cfg *Config) (gfx.Engine, error) {
	switch cfg.Engine {
	case EngineSDL:
		return sdl.NewEngine(&cfg.Graphics)
	default:
		return nil, ErrUnknownEngine
	}
}
