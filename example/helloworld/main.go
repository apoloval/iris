package main

import (
	"fmt"
	"image"
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

		app.BeginLayoutH(250)
		app.BeginLayoutV(200)
		app.BeginLayoutV(5)

		if app.Label(1, "Hello World!", karen.FontColor(col)) {
			col = gfx.ColorRed
		} else {
			col = gfx.ColorWhite
		}

		stats := app.Stats()
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		app.Label(2, "Performance statistics:")

		app.BeginLayoutV(5)
		app.BeginLayoutH(5)
		app.Label(3, "Frames per second :", karen.Expand(image.Pt(300, 0)), karen.Align(gfx.AlignRight))
		app.Label(4, fmt.Sprintf("%.2f", stats.FramesPerSecond))
		app.EndLayout()

		app.BeginLayoutH(5)
		app.Label(5, "Frame render time :", karen.Expand(image.Pt(300, 0)), karen.Align(gfx.AlignRight))
		app.Label(6, fmt.Sprintf("%v", stats.FrameRenderTime))
		app.EndLayout()

		app.BeginLayoutH(5)
		app.Label(7, "Allocated memory :", karen.Expand(image.Pt(300, 0)), karen.Align(gfx.AlignRight))
		app.Label(8, bytesForHuman(mem.Alloc))
		app.EndLayout()
		app.EndLayout()

		app.EndLayout()
		app.EndLayout()
		app.EndLayout()

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
