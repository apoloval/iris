package gfx

// Event is a graphics engine event
type Event interface{}

// EventQuit is the event of quitting the application
type EventQuit struct{}

// Engine is a graphics engine
type Engine interface {
	// Canvas returns the graphics canvas for this engine
	Canvas() Canvas

	// RenderText renders the given text
	RenderText(text string, params RenderTextParams) (RenderedText, error)

	// WaitEvent waits until a graphics engine event is produced, and returns it
	WaitEvent() (Event, error)
}
