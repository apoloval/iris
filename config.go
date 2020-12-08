package karen

import "github.com/apoloval/karen/gfx"

// EngineType is a config setting that defines the graphics engine type
type EngineType string

// Engine types
const (
	EngineSDL EngineType = "sdl"
)

// Config is the application config
type Config struct {
	Engine   EngineType
	Graphics gfx.Config
}

func defaultConfig() *Config {
	return &Config{
		Engine:   EngineSDL,
		Graphics: gfx.DefaultConfig(),
	}
}
