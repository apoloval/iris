package gfx

// Config is the graphics engine configuration
type Config struct {
	ScreenWidth  int
	ScreenHeight int
	WindowTitle  string
}

var defaultConfig = Config{
	ScreenWidth:  800,
	ScreenHeight: 600,
	WindowTitle:  "Karen application",
}

// DefaultConfig generates a config struct with default values
func DefaultConfig() Config {
	return defaultConfig
}
