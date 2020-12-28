package iris

import (
	"errors"
	"time"

	"github.com/apoloval/iris/gfx"
	"github.com/apoloval/iris/internal/app"
	"github.com/apoloval/iris/internal/sdl"
	"github.com/apoloval/iris/internal/widget"
)

// ErrUnknownEngine is an error returned when an unknown engine is specified
var ErrUnknownEngine = errors.New("unknown GFX engine")

// App is the Iris application object.
type App struct {
	config *AppConfig
	state  *app.State
}

// NewApp instantiates a new application
func NewApp(opts ...AppOption) (*App, error) {
	cfg := DefaultConfig()
	if err := cfg.Apply(opts); err != nil {
		return nil, err
	}

	engine, err := newEngine(cfg)
	if err != nil {
		return nil, err
	}

	state := app.NewState(engine)

	app := &App{
		config: cfg,
		state:  state,
	}
	return app, nil
}

// BeginFrame indicates to the app the beginning of a new frame
func (a *App) BeginFrame() {
	a.state.BeginFrame()
}

// EndFrame indicates to the app the end of the current frame
// It returns true if the application was requested to quit
func (a *App) EndFrame() bool {
	a.state.EndFrame()
	return a.state.IO.Quit
}

// BeginLayoutH begins a new horizontal layout
func (a *App) BeginLayoutH(opts ...LayoutOption) {
	a.state.BeginLayoutH(a.applyLayoutOpts(opts))
}

// BeginLayoutV begins a new vertical layout
func (a *App) BeginLayoutV(opts ...LayoutOption) {
	a.state.BeginLayoutV(a.applyLayoutOpts(opts))
}

// EndLayout ends the current layout
func (a *App) EndLayout() {
	a.state.EndLayout()
}

// Stats returns the application performance statistics
func (a *App) Stats() Stats {
	frt := a.state.LastFrameDuration
	fps := 1.0 / frt.Seconds()
	return Stats{
		FramesPerSecond: fps,
		FrameRenderTime: frt,
	}
}

// Label places a new label widget.
// Returns true if the label is mouse focused.
func (a *App) Label(wid uint, text string, opts ...WidgetOption) bool {
	a.applyWidgetOpts(opts)
	return widget.Label(a.state, wid, text)
}

func (a *App) applyWidgetOpts(opts []WidgetOption) {
	for _, opt := range opts {
		opt(&a.state.DrawProps)
	}
}

func (a *App) applyLayoutOpts(opts []LayoutOption) app.LayoutProps {
	var props app.LayoutProps
	for _, opt := range opts {
		opt(&props)
	}
	return props
}

func newEngine(cfg *AppConfig) (gfx.Engine, error) {
	switch cfg.Engine {
	case EngineSDL:
		return sdl.NewEngine(&cfg.Graphics)
	default:
		return nil, ErrUnknownEngine
	}
}

// Stats are application performance statistics
type Stats struct {
	FramesPerSecond float64
	FrameRenderTime time.Duration
}
