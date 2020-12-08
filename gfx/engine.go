package gfx

// Event is a graphics engine event
type Event interface{}

// EventQuit is the event of quitting the application
type EventQuit struct{}

// Engine is a graphics engine
type Engine interface {
	Canvas() Canvas
	WaitEvent() (Event, error)
}
