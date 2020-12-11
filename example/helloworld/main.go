package main

import (
	"fmt"
	"runtime"

	"github.com/apoloval/karen"
	"github.com/apoloval/karen/gfx"
)

func main() {
	app, err := karen.NewApp()
	if err != nil {
		panic(err)
	}

	col := gfx.ColorWhite
	for {
		app.BeginFrame()
		if app.Label(1, "Hello World!", karen.FontColor(col)) {
			col = gfx.ColorRed
		} else {
			col = gfx.ColorWhite
		}

		stats := app.Stats()
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		app.Label(2, "Performance statistics:")
		app.Label(3, fmt.Sprintf("   Frames per second : %.2f", stats.FramesPerSecond))
		app.Label(4, fmt.Sprintf("   Frame render time : %v", stats.FrameRenderTime))
		app.Label(5, fmt.Sprintf("   Allocated memory  : %s", bytesForHuman(mem.Alloc)))

		if app.EndFrame() {
			break
		}
	}
}

func bytesForHuman(bytes uint64) string {
	if bytes < 10*1024 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 10*1024*1024 {
		return fmt.Sprintf("%d KB", bytes/1024)
	} else if bytes < 10*1024*1024*1024 {
		return fmt.Sprintf("%d MB", bytes/(1024*1024))
	} else if bytes < 10*1024*1024*1024*1024 {
		return fmt.Sprintf("%d GB", bytes/(1024*1024*1024))
	}
	return fmt.Sprintf("%d B", bytes)
}
