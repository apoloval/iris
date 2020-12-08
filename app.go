package karen

import "errors"

// ErrAppAlreadyExists indicates an error when application already exists
var ErrAppAlreadyExists = errors.New("Application already exists")

// App is the Karen application object.
type App struct {
	config Config
}

// NewApp instantiates a new application
func NewApp(opts ...Option) (*App, error) {
	var cfg Config
	if err := applyOptions(&cfg, opts); err != nil {
		return nil, err
	}
	app := &App{
		config: cfg,
	}
	return app, nil
}
