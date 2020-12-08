package karen

import (
	"errors"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/sdl"
)

// ErrUnknownEngine is an error returned when an unknown engine is specified
var ErrUnknownEngine = errors.New("unknown GFX engine")

// App is the Karen application object.
type App struct {
	config   *Config
	graphics gfx.Engine
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
func (a *App) NewScene() *Scene {
	return &Scene{}
}

// Run this application until closed or fails
func (a *App) Run() error {
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
