package iris

import "github.com/apoloval/iris/gfx"

// EngineType is a config setting that defines the graphics engine type
type EngineType string

// Engine types
const (
	EngineSDL EngineType = "sdl"
)

// AppConfig is the application config
type AppConfig struct {
	Engine   EngineType
	Graphics gfx.Config
}

// DefaultConfig returns the default app config
func DefaultConfig() *AppConfig {
	return &AppConfig{
		Engine:   EngineSDL,
		Graphics: gfx.DefaultConfig(),
	}
}

// Apply the given app options to this config
func (c *AppConfig) Apply(opts []AppOption) error {
	for _, o := range opts {
		if err := o(c); err != nil {
			return err
		}
	}
	return nil
}
