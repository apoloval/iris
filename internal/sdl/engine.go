package sdl

import (
	"errors"
	"fmt"
	"image"
	"runtime"

	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/internal/io"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Engine is the Engine graphics engine
type Engine struct {
	window *sdl.Window
	screen *sdl.Surface
	fonts  map[string]*ttf.Font
	dpi    *dpi
}

// NewEngine instantiates the SDL graphics engine
func NewEngine(cfg *gfx.Config) (*Engine, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if err := ttf.Init(); err != nil {
		return nil, err
	}

	dpi, err := initDPI()
	if err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(
		cfg.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(dpi.scaleX(cfg.ScreenWidth)),
		int32(dpi.scaleY(cfg.ScreenHeight)),
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
		dpi:    dpi,
	}
	return engine, nil
}

// BeginFrame begins a new screen frame
func (e *Engine) BeginFrame() {
	e.screen.FillRect(nil, 0)
}

// EndFrame ends the current screen frame
func (e *Engine) EndFrame() {
	if err := e.window.UpdateSurface(); err != nil {
		panic(err)
	}
}

// ScreenDims returns the dimensions of the screen
func (e *Engine) ScreenDims() image.Point {
	return e.screen.Bounds().Size()
}

// TextDims calculates the dimensions of the given text
func (e *Engine) TextDims(text string, params gfx.RenderTextParams) (image.Point, error) {
	font, err := e.font(params.Font, params.Size)
	if err != nil {
		return image.Point{}, err
	}

	w, h, err := font.SizeUTF8(text)
	if err != nil {
		return image.Point{}, err
	}
	return image.Pt(w, h), nil
}

// Apply the given render actions
func (e *Engine) Apply(actions []gfx.DrawAction) error {
	for _, act := range actions {
		var err error
		switch a := act.(type) {
		case gfx.DrawTexture:
			err = e.applyDrawTexture(a)
		case gfx.DrawText:
			err = e.applyDrawText(a)
		default:
			err = errors.New("unknown draw action received by the engine")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) applyDrawText(act gfx.DrawText) error {
	if err := act.Params.Validate(); err != nil {
		return err
	}

	font, err := e.font(act.Params.Font, act.Params.Size)
	if err != nil {
		return err
	}

	surface, err := font.RenderUTF8Blended(act.Text, sdl.Color(act.Params.Color))
	if err != nil {
		return err
	}
	defer surface.Free()

	srcSize := gfx.AsRectSize(image.Pt(int(surface.W), int(surface.H)))
	destSize := gfx.AsRectSize(act.Dest.Size())
	srcRect := srcSize.Intersect(destSize)

	return e.drawSurface(
		srcRect,
		act.Dest,
		surface,
	)
}

func (e *Engine) applyDrawTexture(act gfx.DrawTexture) error {
	switch t := act.Texture.(type) {
	case *texture:
		return e.drawSurface(act.Src, act.Dest, t.surface)
	default:
		return errors.New("invalid texture type received by the engine")
	}
}

func (e *Engine) drawSurface(src, dest image.Rectangle, sur *sdl.Surface) error {
	srcRect := ToSDLRect(src)
	dstRect := ToSDLRect(dest)

	return sur.Blit(&srcRect, e.screen, &dstRect)
}

// PollEvents polls events from SDL engine and updates the given IO state
func (e *Engine) PollEvents(s *io.State) error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.MouseMotionEvent:
			s.MousePos.X = int(ev.X)
			s.MousePos.Y = int(ev.Y)
		case *sdl.QuitEvent:
			s.Quit = true
			return nil
		}
	}
	return nil
}

func (e *Engine) font(fontType gfx.TextFontType, fontSize gfx.TextFontSize) (*ttf.Font, error) {
	key := fmt.Sprintf("%s:%d", fontType, fontSize)

	if f, ok := e.fonts[key]; ok {
		return f, nil
	}

	f, err := ttf.OpenFont(e.fontPath(fontType), e.dpi.scaleY(int(fontSize)))
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

type texture struct {
	surface *sdl.Surface
}

func (t *texture) Size() image.Point {
	return image.Point{
		X: int(t.surface.W),
		Y: int(t.surface.H),
	}
}
